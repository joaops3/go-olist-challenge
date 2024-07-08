package controllers

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/src/api/services"
	"github.com/joaops3/go-olist-challenge/src/helpers"
)

type MovieController struct {
	MovieService services.MovieService
}

func InitMovieController(movieService services.MovieService) *MovieController{
	controller := &MovieController{MovieService: movieService}
	return controller
}

func (c *MovieController) UploadMovieCsvChunks(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	
	if err != nil {
		helpers.SendError(ctx, 400, "Erro ao enviar imagem")
		return
	}
	
	src, err := file.Open()
	if err != nil {
		helpers.SendError(ctx, 500, "Erro ao abrir imagem")
		return 
	}

	defer src.Close()

	reader := bufio.NewReader(src)
    buffer := new(bytes.Buffer)
    header, err := reader.ReadBytes('\n')
    if err != nil && err != io.EOF {
        helpers.SendError(ctx, 500, "Erro ao ler arquivo")
        return
    }
    for {
        line, err := reader.ReadBytes('\n')
        if err != nil && err != io.EOF {
            helpers.SendError(ctx, 500, "Erro ao ler arquivo")
            return
        }
        buffer.Write(header)
        buffer.Write(line)
        if err == io.EOF {
            break
        }

        if buffer.Len() > 1024 { 
            if _, err = c.MovieService.UploadCsvChunks(buffer.Bytes()); err != nil {
                helpers.SendError(ctx, 500, err.Error())
                return
            }
            buffer.Reset()
        }
    }

    
    if buffer.Len() > 0 {
        chunk := append(header, buffer.Bytes()...)
        if _, err = c.MovieService.UploadCsvChunks(chunk); err != nil {
            helpers.SendError(ctx, 500, err.Error())
            return
        }
    }
	helpers.SendSuccess(ctx, "success", true)
	return
}

func (c *MovieController) UploadMovieCsv(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	
	if err != nil {
		helpers.SendError(ctx, 400, "Erro ao enviar imagem")
		return
	}
	
	src, err := file.Open()
	if err != nil {
		helpers.SendError(ctx, 500, "Erro ao abrir imagem")
		return 
	}

	defer src.Close()

	reader := bufio.NewReader(src)
    csvReader := csv.NewReader(reader)
    _, err = csvReader.Read()
    if err != nil && err != io.EOF {
        helpers.SendError(ctx, 500, "Erro ao ler arquivo")
        return
    }
    for {
        line, err := csvReader.Read()
        if err != nil && err != io.EOF {
            helpers.SendError(ctx, 500, "Erro ao ler arquivo")
            return
        }
     
        if err == io.EOF {
            break
        }

       go  c.MovieService.UploadCsv(line)
       if err != nil {
            helpers.SendError(ctx, 500, err.Error())
            return
        }
    }

	helpers.SendSuccess(ctx, "success", true)
	return
}



func (c *MovieController) GetPaginated(ctx *gin.Context) {
   data, err :=  c.MovieService.GetPaginated()

   if err != nil {
     helpers.SendError(ctx, 400, err.Error())
     return
   }

   helpers.SendSuccess(ctx, "success", data)
   return
}

func (c *MovieController) GetOne(ctx *gin.Context) {

}

func (c *MovieController) Post(ctx *gin.Context) {

}

func (c *MovieController) Update(ctx *gin.Context) {

}
func (c *MovieController) Delete(ctx *gin.Context) {

}