package SearchableDeveloper

type DeveloperSearchableFields struct {
  Name string `form:"name" json:"name,omitempty"`
  Designation string `form:"designation" json:"designation,omitempty"`
  LookingForJob string `form:"actively_job_searching" json:"actively_job_searching,omitempty"`
  MaxExperience string `form:"max_exp" json:"max_exp,omitempty"`
  MinExperience string `form:"min_exp" json:"min_exp,omitempty"`
  Skills string `form:"skills" json:"skills,omitempty"`
}
