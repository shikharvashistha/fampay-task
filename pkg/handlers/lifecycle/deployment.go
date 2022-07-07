package lifecycle

import (
	"context"
	"math"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/shikharvashistha/fampay/pkg/store"
	"github.com/shikharvashistha/fampay/pkg/store/relational/models.go"
	"github.com/shikharvashistha/fampay/pkg/types"
	"github.com/shikharvashistha/fampay/pkg/utils"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/robfig/cron.v2"
)

func NewDeploymentSvc(store store.Store, log *utils.Logger) Deployment {
	return &deploy{
		Store:  store,
		Logger: log,
	}
}

type deploy struct {
	*utils.Logger
	store.Store
}

var c = cron.New()

// GetData implements Deployment
func (svc *deploy) GetData(genere, page, limit string) (types.Response, error) {
	videoRecords, err := svc.Store.RL().Videos().List(&models.Videos{})
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to get deployment records")
		return types.Response{}, err
	}
	pageNo, err := strconv.Atoi(page)
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to convert page to int")
		return types.Response{}, err
	}
	pageSize, err := strconv.Atoi(limit)
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to convert limit to int")
		return types.Response{}, err
	}
	// Pagination
	start := pageNo * pageSize
	end := start + pageSize
	if end > len(videoRecords) {
		end = len(videoRecords)
	}
	if len(videoRecords) == 0 {
		return types.Response{}, nil
	}
	videoRecords = videoRecords[start:end]
	var Videos []types.VideoDataResponse
	for _, deploymentRecord := range videoRecords {
		if deploymentRecord.Genere == genere {
			Videos = append(Videos, types.VideoDataResponse{
				VideoTitle:    deploymentRecord.VideoTitle,
				Description:   deploymentRecord.Description,
				Publishing:    deploymentRecord.Publishing,
				ThumnailsURLs: deploymentRecord.ThumnailsURLs,
			})
		}
	}
	return types.Response{
		VideoData: Videos,
		PageNo:    pageNo,
	}, nil
}

// SearchData implements Deployment
func (svc *deploy) SearchData(title string, description string) (types.Response, error) {
	videoRecords, err := svc.Store.RL().Videos().List(&models.Videos{})
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to get deployment records")
		return types.Response{}, err
	}
	title = strings.ToLower(title)
	description = strings.ToLower(description)
	var Videos []types.VideoDataResponse
	for _, deploymentRecord := range videoRecords {
		if strings.Contains(deploymentRecord.VideoTitle, title) || strings.Contains(deploymentRecord.Description, description) {
			Videos = append(Videos, types.VideoDataResponse{
				VideoTitle:    deploymentRecord.VideoTitle,
				Description:   deploymentRecord.Description,
				Publishing:    deploymentRecord.Publishing,
				ThumnailsURLs: deploymentRecord.ThumnailsURLs,
			})
		}
	}
	return types.Response{
		VideoData: Videos,
	}, nil
}

// CronSchedule implements Deployment
func (svc *deploy) CronSchedule(request types.CronRequest) (types.CronResponse, error) {
	var videoID string
	c.AddFunc("@every "+request.Interval, func() {
		for _, key := range request.APIKeys {
			var client *youtube.Service
			var err error
			if client, err = youtube.NewService(context.TODO(), option.WithAPIKey(key)); err != nil {
				svc.Logger.WithError(utils.Platform, err).Error("API key is invalid")
				continue
			}
			call := client.Search.List([]string{"id,snippet"}).
				Q(request.Query).
				MaxResults(25).
				Order("date").
				Type("video").
				PublishedAfter("2020-01-01T00:00:00Z")

			if response, err := call.Do(); err != nil {
				svc.Logger.WithError(utils.Platform, err).Error("failed to get video metadata from youtube")
				return
			} else {
				for _, item := range response.Items {
					var thumbnail string
					videoID = item.Id.VideoId
					thumbnail += item.Snippet.Thumbnails.Default.Url + ","
					deploymentRecord := models.Videos{
						Model: models.Model{
							ID: item.Id.VideoId,
						},
						VideoTitle:    item.Snippet.Title,
						Description:   item.Snippet.Description,
						Publishing:    item.Snippet.PublishedAt,
						Genere:        item.Kind,
						ThumnailsURLs: thumbnail,
					}
					err := svc.Store.RL().Videos().Create(&deploymentRecord)
					if err != nil {
						svc.Logger.WithError(utils.Platform, err).Error("failed to create deployment record")
						return
					}
				}
			}

		}
	})
	c.Start()
	var CronID float64
	cronRecords, err := svc.Store.RL().Cron().List(&models.Cron{})
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to get cron records")
		return types.CronResponse{}, err
	}
	for _, cronRecord := range cronRecords {
		value, err := strconv.Atoi(cronRecord.CronID)
		if err != nil {
			svc.Logger.WithError(utils.Platform, err).Error("failed to convert cronID to int")
			return types.CronResponse{}, err
		}
		cronID := float64(value)
		CronID = math.Max(cronID, CronID)
	}
	if videoID == "" {
		videoID = uuid.New().String()
	}
	cron_id := strconv.Itoa(int(CronID + 1))
	cronRecord := models.Cron{
		CronModel: models.CronModel{
			Model: models.Model{
				ID: videoID,
			},
			CronID: cron_id,
		},
	}
	err = svc.Store.RL().Cron().Create(&cronRecord)
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to create cron record")
		return types.CronResponse{}, err
	}
	return types.CronResponse{ID: cron_id}, nil
}
func (svc *deploy) CronDelete(request types.CronDeleteRequest) error {
	cronID, err := strconv.Atoi(request.ID)
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to convert cron ID to int")
		return err
	}
	cronModel := &models.Cron{
		CronModel: models.CronModel{
			CronID: request.ID,
		},
	}
	err = svc.Store.RL().Cron().Delete(cronModel)
	if err != nil {
		svc.Logger.WithError(utils.Platform, err).Error("failed to delete cron record")
		return err
	}
	cronEntryID := cron.EntryID(cronID)
	c.Remove(cronEntryID)
	return nil
}
