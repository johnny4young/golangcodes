package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golangcodes/commons"
	"github.com/golangcodes/configuration"
	"github.com/golangcodes/models"
)

// VoteRegister controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currenVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el voto a registrar: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? and user_id = ?", vote.CommentID, vote.UserID).First(&currenVote)

	// si no existe
	if currenVote.ID == 0 {
		db.Create(&vote)
		err := updateCommentVote(vote.CommentID, vote.Value, false)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w, m)
		return
	} else if currenVote.Value != vote.Value {
		currenVote.Value = vote.Value
		db.Save(&currenVote)
		err := updateCommentVote(vote.CommentID, vote.Value, true)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto actualizado"
		m.Code = http.StatusOK
		commons.DisplayMessage(w, m)
	}
	m.Message = "Este voto ya està registrado"
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}

// updateCommentVote actualiza la cantidad de votos
func updateCommentVote(commentID uint, vote bool, isUpdate bool) (err error) {
	comment := models.Comment{}

	db := configuration.GetConnection()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected
	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate {
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontró un registro de comentario para asignarle un comentario")
	}
	return
}
