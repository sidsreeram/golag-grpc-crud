package pkg
import(
	"github.com/msproject2/internal/models"
)
type UserRepository interface{
 Create(user models.User) (string, error)
	Get(id string) (models.User, error)
	Update(user models.User) error 
	Delete(id string) error
	GetByEmail(email string) (models.User, error)
}
type UseUsecase interface{

    Create(user models.User) (string, error)
	Get(id string) (models.User, error)
	Update(updateUser models.User) error
	Delete(id string) error
}