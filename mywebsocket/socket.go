package mywebsocket

/*func  Socket(server *socketio.Server) *socketio.Server {

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		return nil
	})

	server.OnEvent("/", "joinroom", func(s socketio.Conn, roomnum string) {
		if server.RoomLen("", roomnum) >= 2 {
			s.Emit("joinroom", "2인 이상 집합금지")
			return
		}
		log.Println("joinroom roomnumber : " + roomnum + "id : " + s.ID())
		s.Join(roomnum)
		server.BroadcastToRoom("/", roomnum, "joinroom", s.ID()+" is join "+roomnum+" room")
	})

	server.OnEvent("/", "leaveroom", func(s socketio.Conn, roomnum string) {
		s.Leave(roomnum)
		server.BroadcastToRoom("/", roomnum, "leaveroom", s.ID()+" is leave "+roomnum+" room")
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg models.Message) {
		s.SetContext(msg)
		err := database.DB.Model(&models.Message{}).Create(msg).Error
		if err != nil {
			server.BroadcastToRoom("/", msg.NumRoom, "error", msg.UserId+": "+"message create failed")
		}
		server.BroadcastToRoom("/", msg.NumRoom, "msg", msg.UserId+": "+msg.Content)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed " + s.ID() + " " + reason)
		s.Close()
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		log.Println("error : ", err)
	})

	return server
}*/
