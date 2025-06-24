package repository

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/blockloop/scan"
	"github.com/textures1245/payso-check-slip-backend/model"

	log "github.com/sirupsen/logrus"
)

func GetAllRooms() ([]model.SureSureRoom, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureRoom{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ROOM_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureRoom{}, err
	}

	var rooms []model.SureSureRoom
	err = scan.Rows(&rooms, rows)

	defer rows.Close()
	log.Infof("rooms: %d", len(rooms))
	return rooms, nil
}

func GetRoomByID(id int) ([]model.SureSureRoom, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureRoom{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ROOM_GET_BYUSERID, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return []model.SureSureRoom{}, err
	}
	var room []model.SureSureRoom
	err = scan.Rows(&room, rows)
	defer rows.Close()

	return room, nil
}

func CreateRoom(room model.SureSureRoom) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	// Build query dynamically
	query := "INSERT INTO SureSureRoom ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1
	if room.UserID != 0 {
		query += "UserID, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.UserID))
		counter++
	}

	if room.LineGroupID != "" {
		query += "LineGroupID, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.LineGroupID))
		counter++
	}

	if room.RoomName != "" {
		query += "RoomName, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.RoomName))
		counter++
	}

	if room.QRToken != "" {
		query += "QRToken, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.QRToken))
		counter++
	}

	if room.QuotaUsed != 0 {
		query += "QuotaUsed, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.QuotaUsed))
		counter++
	}

	if room.MinRecieve != 0 {
		query += "MinRecieve, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.MinRecieve))
		counter++
	}

	if room.ShowTransferor {
		query += "ShowTransferor, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ShowTransferor))
		counter++
	}

	if room.ShowRecipient {
		query += "ShowRecipient, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ShowRecipient))
		counter++
	}

	if room.ListBank != "" {
		query += "ListBank, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ListBank))
		counter++
	}

	if room.CreatedDate != "" {
		query += "CreatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.CreatedDate))
		counter++
	}

	if room.UpdatedDate != "" {
		query += "UpdatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.UpdatedDate))
		counter++
	}

	query = query[:len(query)-2] + ") "
	values = values[:len(values)-2] + ")"
	finalQuery := query + " " + values + ";select ID = convert(bigint, SCOPE_IDENTITY())"

	log.Infof("finalQuery: %v", finalQuery)
	result, err := conn.ExecContext(ctx, finalQuery, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return 0, err
	}
	log.Infof("result: %v", result)
	// Retrieve the last inserted ID
	// lastInsertedID, err := result.LastInsertId()
	// if err != nil {
	// 	log.Errorf("Error retrieving last insert ID: %v", err)
	// 	return 0, err
	// }

	return 0, nil
}

func UpdateRoom(room model.SureSureRoom) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// Initialize query parts
	query := "UPDATE SureSureRoom SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if room.UserID != 0 {
		query += fmt.Sprintf("UserID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.UserID))
		counter++
	}

	if room.LineGroupID != "" {
		query += fmt.Sprintf("LineGroupID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.LineGroupID))
		counter++
	}

	if room.RoomName != "" {
		query += fmt.Sprintf("RoomName = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.RoomName))
		counter++
	}

	if room.QRToken != "" {
		query += fmt.Sprintf("QRToken = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.QRToken))
		counter++
	}

	if room.QuotaUsed != 0 {
		query += fmt.Sprintf("QuotaUsed = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.QuotaUsed))
		counter++
	}

	if room.MinRecieve != 0 {
		query += fmt.Sprintf("MinRecieve = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.MinRecieve))
		counter++
	}

	if room.ShowTransferor != false {
		query += fmt.Sprintf("ShowTransferor = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ShowTransferor))
		counter++
	}

	if room.ShowRecipient != false {
		query += fmt.Sprintf("ShowRecipient = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ShowRecipient))
		counter++
	}

	if room.ListBank != "" {
		query += fmt.Sprintf("ListBank = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ListBank))
		counter++
	}

	if room.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.UpdatedDate))
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = @p" + fmt.Sprintf("%d", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), room.ID))

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteRoom(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	log.Infof("id: %d", id)
	rows, err := conn.QueryContext(ctx, model.SQL_ROOM_DELETE, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}

func HowTo(id int, user_id string) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// ROOM
	rows, err := conn.QueryContext(ctx, model.SQL_ROOM_GET_BYID, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	var room model.SureSureRoom
	err = scan.Row(&room, rows)
	defer rows.Close()

	// USER
	rows, err = conn.QueryContext(ctx, model.SQL_USER_GET_BYID, sql.Named("ID", room.UserID))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	var user model.SureSureUser
	err = scan.Row(&user, rows)
	defer rows.Close()

	// LINE Messaging API endpoint for push messages
	url := "https://api.line.me/v2/bot/message/push"

	status := ""
	statusColor := "#667085"

	if room.LineGroupID == "" {
		status = "ยังไม่เชื่อมต่อ"
		statusColor = "#D92D20"
	} else {
		status = "เชื่อมต่อแล้ว"
		statusColor = "#1DB446"
	}
	if user.StoreName == "" {
		user.StoreName = "-"
	}
	if room.RoomName == "" {
		room.RoomName = "-"
	}
	// Prepare the payload
	payload := map[string]interface{}{
		"to": user_id, // User ID to whom the message will be sent
		"messages": []map[string]interface{}{
			{
				"type":    "flex",
				"altText": "กรุณาเชื่อมต่อสาขา",
				"contents": map[string]interface{}{
					"type": "bubble",
					"body": map[string]interface{}{
						"type":   "box",
						"layout": "vertical",
						"contents": []map[string]interface{}{
							{
								"type":   "box",
								"layout": "baseline",
								"contents": []map[string]interface{}{
									{"type": "icon", "url": "https://img2.pic.in.th/pic/Featured-icon-9525f708b1582d30e.png", "size": "lg"},
									{"type": "text", "text": "กรุณาเชื่อมต่อสาขา", "weight": "bold", "size": "xl", "color": "#D92D20", "margin": "md"},
								},
							},
							{"type": "text", "text": "1. สร้าง LINE Group  และเชิญ SureSure", "wrap": true, "color": "#667085", "size": "sm", "margin": "lg"},
							{"type": "text", "text": "2. ส่ง Code สำหรับเชื่อมต่อ ใน LINE Group", "wrap": true, "color": "#667085", "size": "sm", "margin": "sm"},
							{"type": "text", "text": "3. เชื่อมต่อสำเร็จ ตรวจสลิปได้ทันที", "wrap": true, "color": "#667085", "size": "sm", "margin": "sm"},
							{"type": "separator", "margin": "lg"},
							{
								"type":   "box",
								"layout": "horizontal",
								"margin": "md",
								"contents": []map[string]interface{}{
									{"type": "text", "text": "ร้าน", "color": "#667085", "size": "sm", "flex": 2},
									{"type": "text", "text": user.StoreName, "size": "sm", "flex": 2},
								},
							},
							{
								"type":   "box",
								"layout": "horizontal",
								"margin": "sm",
								"contents": []map[string]interface{}{
									{"type": "text", "text": "สาขา", "color": "#667085", "size": "sm", "flex": 2},
									{"type": "text", "text": room.RoomName, "size": "sm", "flex": 2},
								},
							},
							{
								"type":   "box",
								"layout": "horizontal",
								"margin": "sm",
								"contents": []map[string]interface{}{
									{"type": "text", "text": "เลขที่อ้างอิงสาขา", "color": "#667085", "size": "sm", "flex": 2},
									{"type": "text", "text": "#" + fmt.Sprintf("%05d", room.ID), "size": "sm", "flex": 2},
								},
							},
							{
								"type":   "box",
								"layout": "horizontal",
								"margin": "sm",
								"contents": []map[string]interface{}{
									{"type": "text", "text": "สถานะ", "color": "#667085", "size": "sm", "flex": 2},
									{"type": "text", "text": status, "color": statusColor, "size": "sm", "flex": 2},
								},
							},
						},
					},
				},
			},
		},
	}

	// Convert payload to JSON
	jsonData, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Errorf("Error marshaling payload: %v", err)
		return err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorf("Error creating request: %v", err)
		return err
	}
	accessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	// Set headers for the LINE Messaging API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Errorf("Error sending message, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}
