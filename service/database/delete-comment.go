package database

func (db *appdbimpl) DeleteComment(comment_id int) (string, error) {
	_, err := db.c.Exec("DELETE FROM comments WHERE comment_id=?", comment_id)
	if err != nil {
		return "", err
	}
	return "commento rimosso correttamente", nil
}
