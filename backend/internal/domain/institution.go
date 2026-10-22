package domain

type Institution struct {
 ID string `json:"id"`
 Name string `json:"name"`
 City string `json:"city"`
 Kind string `json:"kind"`
 Website string `json:"website"`
 Description string `json:"description"`
 Tags []string `json:"tags"`
}
