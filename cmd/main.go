package main

import
(
	"github.com/BrakovIgor19/backend-trainee-assignment-2023-"
	"log"
)

func main()
{
	srv := new(segmentationService.Server)

	if err := srv.Run("8080"); err != nil
	{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}