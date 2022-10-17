package worker

import (
	"github.com/matthieurobert/go-api-template/entity"
	"github.com/matthieurobert/go-api-template/errors"
	validator "gopkg.in/validator.v2"

	pg "github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

// IdeaWorker TO DO
type IdeaWorker struct {
	DB  *pg.DB
	Log *logrus.Logger
}

// GetIdeas : Get All Ideas
func (w IdeaWorker) GetIdeas() ([]entity.Idea, *errors.HandleError) {
	w.Log.Debug("Begin worker GetIdeas")

	ideas, handleError := w.RetrieveIdeas()
	if handleError != nil {
		return nil, handleError
	}

	w.Log.Debug("No errors worker GetIdeas.")
	return ideas, nil
}

// GetIdea : Get Idea with one ID
func (w IdeaWorker) GetIdea(ideaID int) (*entity.Idea, *errors.HandleError) {
	w.Log.Debug("Begin worker GetIdea")

	idea, handleError := w.RetrieveIdeaByID(ideaID)
	if handleError != nil {
		return nil, handleError
	}

	w.Log.Debug("No errors worker GetIdea.")
	return idea, nil
}

// PostIdea : Post new idea
func (w IdeaWorker) PostIdea(idea entity.Idea) (*int, *errors.HandleError) {
	w.Log.Debug("Begin worker PostIdea")

	handleError := w.ValidNewIdea(idea)
	if handleError != nil {
		return nil, handleError
	}

	ideaID, handleError := w.InsertIdea(idea)
	if handleError != nil {
		return nil, handleError
	}

	w.Log.Debug("No errors worker PostIdea.")
	return ideaID, nil
}

// UpdateIdea : Update existinge idea
func (w IdeaWorker) UpdateIdea(idea entity.Idea) (*int, *errors.HandleError) {
	w.Log.Debug("Begin worker UpdateIdea")

	handleError := w.ValidNewIdea(idea)
	if handleError != nil {
		return nil, handleError
	}

	ideaID, handleError := w.ModifyIdea(idea)
	if handleError != nil {
		return nil, handleError
	}

	w.Log.Debug("No errors worker UpdateIdea")
	return ideaID, nil
}

// DeleteIdea : Delete idea
func (w IdeaWorker) DeleteIdea(ideaID int) *errors.HandleError {
	w.Log.Debug("Begin worker DeleteIdea")

	handlerError := w.DropIdea(ideaID)
	if handlerError != nil {
		return handlerError
	}

	w.Log.Debug("No errors worker DeleteIdea")
	return nil
}

// BuildRepo : Create repository factory and build repository
func (w IdeaWorker) BuildRepo() *entity.IdeaRepository {
	ideaRepoFactory := &entity.IdeaRepositoryFactory{
		Database: w.DB,
	}
	w.Log.Debug("IdeaRepoFactory created")
	ideaRepo := ideaRepoFactory.Build()
	w.Log.Debug("ideaRepo created")
	return ideaRepo
}

// ValidNewIdea : Validate the idea with tag validator
func (w IdeaWorker) ValidNewIdea(idea entity.Idea) *errors.HandleError {
	err := validator.WithTag("validate").Validate(idea)
	if err != nil {
		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Info("New idea validation failed with tag 'validate'")
		handleError := errors.CreateValidationError()
		return handleError
	}

	err = validator.WithTag("creating").Validate(idea)
	if err != nil {
		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Info("New idea validation failed with tag 'creating'")
		handleError := errors.CreateValidationError()
		return handleError
	}

	w.Log.Debug("New idea is valid")
	return nil
}

// InsertIdea : Insert valid idea into database
func (w IdeaWorker) InsertIdea(idea entity.Idea) (*int, *errors.HandleError) {
	ideaRepo := w.BuildRepo()
	ideaID, err := ideaRepo.PostIdea(idea)

	if err != nil {
		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Errorf("Server error calling idea repository : PostIdea method.")
		handleError := errors.CreateServerError()
		return nil, handleError
	}

	w.Log.WithFields(logrus.Fields{
		"ideaID": ideaID,
	}).Debug("New idea inserted.")
	return &ideaID, nil
}

// ModifyIdea : Update valid idea into database
func (w IdeaWorker) ModifyIdea(idea entity.Idea) (*int, *errors.HandleError) {
	ideaRepo := w.BuildRepo()
	ideaID, err := ideaRepo.UpdateIdea(idea)

	if err != nil {
		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Errorf("Server error calling idea repository : UpdateIdea method.")
		handleError := errors.CreateServerError()
		return nil, handleError
	}

	w.Log.WithFields(logrus.Fields{
		"ideaID": ideaID,
	}).Debug("Idea updated.")
	return &ideaID, nil
}

// RetrieveIdeas : Call repository to get ideas
func (w IdeaWorker) RetrieveIdeas() ([]entity.Idea, *errors.HandleError) {
	ideaRepo := w.BuildRepo()
	ideas, err := ideaRepo.GetAllIdeas()

	if err != nil {
		if err == pg.ErrNoRows {
			w.Log.Debug("Calling idea repository : GetAllIdeas method. No data found")
			return []entity.Idea{}, nil
		}

		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Errorf("Server error calling idea repository : GetAllIdeas method.")
		handleError := errors.CreateServerError()
		return nil, handleError
	}

	w.Log.WithFields(logrus.Fields{
		"ideas": ideas,
	}).Debug("No errors calling idea repository.")
	return ideas, nil
}

// RetrieveIdeaByID : Call repository to get a specific idea
func (w IdeaWorker) RetrieveIdeaByID(ideaID int) (*entity.Idea, *errors.HandleError) {
	ideaRepo := w.BuildRepo()
	idea, err := ideaRepo.GetIdea(ideaID)

	if err != nil {
		if err == pg.ErrNoRows {
			w.Log.Debug("Calling idea repository : GetIdea method. No data found")
			return nil, nil
		}

		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Errorf("Server error calling idea repository : GetIdea method.")
		handleError := errors.CreateServerError()
		return nil, handleError
	}

	w.Log.WithFields(logrus.Fields{
		"idea": idea,
	}).Debug("No errors calling idea repository.")
	return idea, nil
}

// DropIdea : Delete idea into database
func (w IdeaWorker) DropIdea(ideaID int) *errors.HandleError {
	ideaRepo := w.BuildRepo()
	err := ideaRepo.DeleteIdea(ideaID)

	if err != nil {
		if err == pg.ErrNoRows {
			w.Log.Debug("Calling idea repository : DeleteIdea method. No data found")
			return nil
		}

		w.Log.WithFields(logrus.Fields{
			"Err": err,
		}).Errorf("Server error calling idea repository : DeleteIdea method.")
		handleError := errors.CreateServerError()
		return handleError
	}

	w.Log.WithFields(logrus.Fields{
		"ideaID": ideaID,
	}).Debug("No errors calling idea repository")
	return nil
}
