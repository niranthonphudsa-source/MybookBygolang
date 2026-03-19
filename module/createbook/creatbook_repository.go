// connectdata base
package createbook

// import (
// 	"database/sql"
// 	_"github.com/lib/pq"
// )

type CreatbookRepository interface {
	Save(books Books) error
}
