package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/truongtu268/distributePriorityQueue/db/mocks"
	"github.com/truongtu268/distributePriorityQueue/db/query"
	"github.com/truongtu268/distributePriorityQueue/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementations of the query interfaces would be needed here.
// For simplicity, let's assume we have a mock implementation ready.

func TestCreateAdRepo_CreateAd(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdCreateQuery{}
		mockQuery.On("CreateAd", mock.Anything, mock.Anything).Return(nil)
		repo := &CreateAdRepo{
			queries: mockQuery,
		}

		dto := model.AdRequest{
			Title:          "Test Ad",
			Description:    "This is a test ad",
			Genre:          "Test Genre",
			TargetAudience: []string{"Audience1", "Audience2"},
			VisualElements: []string{"Element1", "Element2"},
			CallToAction:   "Click Here",
			Duration:       30,
			Priority:       2,
		}

		err := repo.CreateAd(context.Background(), dto)
		assert.NoError(t, err)
		// Additional assertions can be added here to verify the behavior of the mock
	})

	t.Run("Error case", func(t *testing.T) {
		mockQuery := &mocks.AdCreateQuery{}
		mockQuery.On("CreateAd", mock.Anything, mock.Anything).Return(fmt.Errorf("error something"))
		repo := &CreateAdRepo{
			queries: mockQuery,
		}

		dto := model.AdRequest{
			Title:          "Test Ad",
			Description:    "This is a test ad",
			Genre:          "Test Genre",
			TargetAudience: []string{"Audience1", "Audience2"},
			VisualElements: []string{"Element1", "Element2"},
			CallToAction:   "Click Here",
			Duration:       30,
			Priority:       2,
		}

		err := repo.CreateAd(context.Background(), dto)
		assert.Error(t, err)
		// Additional assertions can be added here to verify the behavior of the mock
	})
}

func TestAdQueueRepo_UpdateAdStatus(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdQueueQuery{}
		mockQuery.On("UpdateAdStatus", mock.Anything, mock.Anything).Return(nil)
		repo := &AdQueueRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		err := repo.UpdateAdStatus(context.Background(), adID)
		assert.NoError(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdQueueQuery{}
		mockQuery.On("UpdateAdStatus", mock.Anything, mock.Anything).Return(fmt.Errorf("error updating status"))
		repo := &AdQueueRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		err := repo.UpdateAdStatus(context.Background(), adID)
		assert.Error(t, err)
	})
}

func TestAdCronjobRepo_ProcessTask(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdStatus", mock.Anything, mock.Anything).Return(nil)
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		err := repo.ProcessTask(context.Background(), adID)
		assert.NoError(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdStatus", mock.Anything, mock.Anything).Return(fmt.Errorf("error processing task"))
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		err := repo.ProcessTask(context.Background(), adID)
		assert.Error(t, err)
	})
}

func TestAdCronjobRepo_AddAdAnalysis(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdAnalysis", mock.Anything, mock.Anything).Return(nil)
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		analysis := model.AdAnalysis{
			// Populate with test data
		}

		err := repo.AddAdAnalysis(context.Background(), adID, analysis)
		assert.NoError(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdAnalysis", mock.Anything, mock.Anything).Return(fmt.Errorf("error adding analysis"))
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		analysis := model.AdAnalysis{
			// Populate with test data
		}

		err := repo.AddAdAnalysis(context.Background(), adID, analysis)
		assert.Error(t, err)
	})
}

func TestAdCronjobRepo_GetAdByID(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		expectedAd := query.Ad{} // Populate with expected data
		mockQuery.On("GetAdByID", mock.Anything, mock.Anything).Return(expectedAd, nil)
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		ad, err := repo.GetAdByID(context.Background(), adID)
		assert.NoError(t, err)
		assert.Equal(t, expectedAd, ad)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("GetAdByID", mock.Anything, mock.Anything).Return(query.Ad{}, fmt.Errorf("error getting ad"))
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		_, err := repo.GetAdByID(context.Background(), adID)
		assert.Error(t, err)
	})
}

func TestAdCronjobRepo_RetryAd(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdRetry", mock.Anything, mock.Anything).Return(nil)
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		retryTime := int32(3)
		err := repo.RetryAd(context.Background(), adID, retryTime)
		assert.NoError(t, err)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdCronjobQuery{}
		mockQuery.On("UpdateAdRetry", mock.Anything, mock.Anything).Return(fmt.Errorf("error retrying ad"))
		repo := &AdCronjobRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		retryTime := int32(3)
		err := repo.RetryAd(context.Background(), adID, retryTime)
		assert.Error(t, err)
	})
}

func TestGetAdRepo_GetAdByID(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		mockQuery := &mocks.AdGetQuery{}
		expectedAd := query.Ad{} // Populate with expected data
		mockQuery.On("GetAdByID", mock.Anything, mock.Anything).Return(expectedAd, nil)
		repo := &GetAdRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		ad, err := repo.GetAdByID(context.Background(), adID)
		assert.NoError(t, err)
		assert.Equal(t, expectedAd, ad)
	})

	t.Run("error case", func(t *testing.T) {
		mockQuery := &mocks.AdGetQuery{}
		mockQuery.On("GetAdByID", mock.Anything, mock.Anything).Return(query.Ad{}, fmt.Errorf("error getting ad"))
		repo := &GetAdRepo{
			queries: mockQuery,
		}

		adID := uuid.New().String()
		_, err := repo.GetAdByID(context.Background(), adID)
		assert.Error(t, err)
	})
}
