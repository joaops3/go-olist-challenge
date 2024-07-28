package controllers

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/api/services"
)

type MovieController struct {
	MovieService services.MovieServiceInterface
}

func InitMovieController(movieService services.MovieServiceInterface) *MovieController {
	controller := &MovieController{MovieService: movieService}
	return controller
}

func (c *MovieController) UploadMovieCsvChunks(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	
	if err != nil {
		sendError(ctx, 400, "Erro ao enviar imagem")
		return
	}
	
	internal, err := file.Open()
	if err != nil {
		sendError(ctx, 500, "Erro ao abrir imagem")
		return 
	}

	defer internal.Close()

	reader := bufio.NewReader(internal)
    buffer := new(bytes.Buffer)
    header, err := reader.ReadBytes('\n')
    if err != nil && err != io.EOF {
       sendError(ctx, 500, "Erro ao ler arquivo")
        return
    }
    for {
        line, err := reader.ReadBytes('\n')
        if err != nil && err != io.EOF {
           sendError(ctx, 500, "Erro ao ler arquivo")
            return
        }
        buffer.Write(header)
        buffer.Write(line)
        if err == io.EOF {
            break
        }

        if buffer.Len() > 1024 { 
            if _, err = c.MovieService.UploadCsvChunks(buffer.Bytes()); err != nil {
               sendError(ctx, 500, err.Error())
                return
            }
            buffer.Reset()
        }
    }

    
    if buffer.Len() > 0 {
        chunk := append(header, buffer.Bytes()...)
        if _, err = c.MovieService.UploadCsvChunks(chunk); err != nil {
           sendError(ctx, 500, err.Error())
            return
        }
    }
	sendSuccess(ctx, "success", true)
	return
}

func (c *MovieController) UploadMovieCsv(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	
	if err != nil {
		sendError(ctx, 400, "Erro ao enviar imagem")
		return
	}
	
	internal, err := file.Open()
	if err != nil {
		sendError(ctx, 500, "Erro ao abrir imagem")
		return 
	}

	defer internal.Close()

	reader := bufio.NewReader(internal)
    csvReader := csv.NewReader(reader)
    _, err = csvReader.Read()
    if err != nil && err != io.EOF {
       sendError(ctx, 500, "Erro ao ler arquivo")
        return
    }
    var wg sync.WaitGroup

    channel := make(chan []string, 10)
    errorsCh := make(chan error)
  
      
    go func() {
        defer close(channel)
        for {
            line, err := csvReader.Read()
            if err != nil {
                if err == io.EOF {
                    return
                }
                errorsCh <- err
                return
            }
            channel <- line
        }
    }()

    go func() {
        defer close(errorsCh)
        for line := range channel {
            wg.Add(1)
            go func(value []string) {
                defer wg.Done()
                _, err := c.MovieService.UploadCsv(value)
                if err != nil {
                    select {
                        case errorsCh <- err:
                    default:
                    }
                }
            }(line)
        
        }
        wg.Wait()
    }()

  
    
    for err := range errorsCh {
		if err != nil {
			sendError(ctx, 422, err.Error())
			return
		}
	}

   sendSuccess(ctx, "success", true)
    return
}



func (c *MovieController) GetPaginated(ctx *gin.Context) {
    

    var paginationDto dtos.PaginationDto

    if err := ctx.BindQuery(&paginationDto); err != nil {
       sendError(ctx, 400, err.Error())
        return
    }


    if err := validator.New().Struct(&paginationDto); err != nil {
       sendError(ctx, 400, err.Error())
        return
    }

   data, err :=  c.MovieService.GetPaginated(&paginationDto)

   if err != nil {
    sendError(ctx, 400, err.Error())
     return
   }

  sendSuccess(ctx, "success", data)
   return
}

func (c *MovieController) GetOne(ctx *gin.Context) {

    id := ctx.Param("id")

    if id == "" {
       sendError(ctx, 400, "id id required")
        return 
    }

     model, err := c.MovieService.GetOne(id)

     if err != nil {
       sendError(ctx, 404, err.Error())
        return
    }

     if model == nil {
        sendError(ctx, 404, "movie not found")
         return
     }

    sendSuccess(ctx, "success", model)
     return

}

func (c *MovieController) Post(ctx *gin.Context) {

    dto := dtos.CreateMovieDto{}

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    err = dto.Validate()
    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.MovieService.Post(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

   sendSuccess(ctx, "success", data)
    return


}

func (c *MovieController) Update(ctx *gin.Context) {

    id := ctx.Param("id")

    if id == "" {
           sendError(ctx, 400, "id id required")
            return 
    }

    var dto dtos.UpdateMovieDto

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.MovieService.Update(id, &dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

   sendSuccess(ctx, "success", data)
    return

}
func (c *MovieController) Delete(ctx *gin.Context) {

    id := ctx.Param("id")

    if id == "" {
       sendError(ctx, 400, "id is required")
        return 
    }

    data, err := c.MovieService.Delete(id)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    
   sendSuccess(ctx, "success", data)
    return

}