package main

import (
	"fmt"
	"log"
	"net"

	dbConfig "github.com/msproject2/internal/db"
	"github.com/msproject2/internal/models"
	interfaces "github.com/msproject2/pkg"
	"github.com/msproject2/pkg/handler"
	repo "github.com/msproject2/pkg/repository"
	usecase "github.com/msproject2/pkg/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	db := dbConfig.DbConn()
	migrations(db)
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}
	grpcServer := grpc.NewServer()
	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)
	log.Fatal(grpcServer.Serve(lis))
}
func initUserServer(db *gorm.DB) interfaces.UseUsecase {
	userRepo := repo.NewUserRepo(db)
	return usecase.NewUsecase(userRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
