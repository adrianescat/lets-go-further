package data

import (
	"time"
)

// Movie struct. It’s crucial to point out here that all the fields in our Movie struct are
// exported (i.e. start with a capital letter), which is necessary for them to be visible to
// Go’s encoding/json package. Any fields which aren’t exported won’t be included
// when encoding a struct to JSON.
type Movie struct {
	ID        int64     // Unique integer ID for the movie
	CreatedAt time.Time // Timestamp for when the movie is added to our database
	Title     string    // Movie title
	Year      int32     // Movie release year
	Runtime   int32     // Movie runtime (in minutes)
	Genres    []string  // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}
