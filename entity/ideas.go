package entity

import (
	"time"

	"github.com/go-pg/pg/v10"
)

// Idea : structure for idea
type Idea struct {
	tableName      struct{}   `sql:"co_idea"`
	ID             int        `json:"id,omitempty" xml:"id,attr" sql:"idea_id,pk"`
	Competition    string     `json:"competition,omitempty" xml:"competition,attr" sql:"idea_competition"`
	ControledField string     `json:"controled_field,omitempty" xml:"controled_field,attr" sql:"idea_controled_field"`
	Costs          string     `json:"costs,omitempty" xml:"costs,attr" sql:"idea_costs"`
	Detail         string     `json:"detail,omitempty" xml:"detail,attr" sql:"idea_detail"`
	EntrustedField string     `json:"entrusted_field,omitempty" xml:"entrusted_field,attr" sql:"idea_entrusted_field"`
	Gain           string     `json:"gain,omitempty" xml:"gain,attr" sql:"idea_gain"`
	ImageID        int        `json:"image_id,omitempty" xml:"imageid,attr" sql:"idea_image_fk_media_id"`
	Market         string     `json:"market,omitempty" xml:"market,attr" sql:"idea_market"`
	MediaID        int        `json:"media_id,omitempty" xml:"mediaid,attr" sql:"idea_video_fk_media_id"`
	Needs          string     `json:"needs,omitempty" xml:"needs,attr" sql:"idea_needs"`
	PublishDate    *time.Time `json:"publish_date,omitempty" xml:"publishdate,attr" sql:"idea_publish_date"`
	Problem        string     `json:"problem,omitempty" xml:"problem,attr" sql:"idea_problem"`
	Summary        string     `json:"summary,omitempty" xml:"summary,attr" sql:"idea_summary"`
	Title          string     `json:"title,omitempty" xml:"title,attr" sql:"idea_title"`
	UserID         string     `json:"user_id,omitempty" xml:"userid,attr" sql:"idea_fk_user_id"`
}

// IdeaRepositoryFactory : structure knowing all parameters to create the repository.
type IdeaRepositoryFactory struct {
	Database *pg.DB
}

// Build : Create the repository
func (f *IdeaRepositoryFactory) Build() *IdeaRepository {
	return &IdeaRepository{
		Database: f.Database,
		Column:   "idea.*",
	}
}

// IdeaRepository : structure knowing all parameters to perform exchange with database for user type.
type IdeaRepository struct {
	Database *pg.DB
	Column   string
}

// GetAllIdeas : get all existing ideas
func (r *IdeaRepository) GetAllIdeas() ([]Idea, error) {
	ideas := []Idea{}
	err := r.Database.Model(&ideas).Select()
	if err != nil {
		return []Idea{}, err
	}
	return ideas, nil
}

// GetIdea : Get a idea by a given ID
func (r *IdeaRepository) GetIdea(id int) (*Idea, error) {
	idea := Idea{ID: id}
	err := r.Database.Model(&idea).Select()
	if err != nil {
		return &Idea{}, err
	}
	return &idea, nil
}

// PostIdea : Post a idea
func (r *IdeaRepository) PostIdea(idea Idea) (int, error) {
	_, err := r.Database.Model(&idea).Insert()

	if err != nil {
		return 0, err
	}
	return idea.ID, nil
}

// UpdateIdea : Update an idea
func (r *IdeaRepository) UpdateIdea(idea Idea) (int, error) {
	_, err := r.Database.Model(&idea).WherePK().Update()

	if err != nil {
		return 0, err
	}

	return idea.ID, nil
}

// DeleteIdea : Delete an idea
func (r *IdeaRepository) DeleteIdea(id int) error {
	idea := Idea{ID: id}
	_, err := r.Database.Model(&idea).Where("id = ?", id).Delete()

	if err != nil {
		return err
	}

	return nil
}
