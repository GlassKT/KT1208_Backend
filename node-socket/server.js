const cors = require("cors");
const express = require('express');
const app = express();
app.use(cors());

var mysql = require('mysql2')
var conn = mysql.createConnection({ // createConnection 메서드로 객체화
});
  
conn.connect();

var sql = 'insert into messages(num_room,content,user_id,createat,name) values(?,?,?,?,?)'

const PORT = 7070;
const server = app.listen(PORT, function() {
    console.log('node Server listen on ' + PORT);
});

const SocketIo = require('socket.io');
const io = SocketIo(server, {path: '/socket.io'});

io.on('connection', function(socket) {

    socket.on('joinroom', (data) => {
        
        socket.join(data.roomnum);
        console.log("join room 작동")

        io.to(data.roomnum).emit("msg", data.name + " is join " + data.roomnum)
    });

    socket.on('leaveroom', (data) => {
        console.log(data)

        console.log("leave room 작동")
        socket.leave(data.roomnum);
        io.to(data.roomnum).emit("msg", data.name + " is leave " + data.roomnum)
    });


    socket.on('msg', (data) => {        // data { roomnum, content, id, name, createat }

        console.log("message is ",data);  

        conn.query(sql,[data.num_room, data.content, data.user_id, data.createat, data.name], function(err, rows, fields){
            console.log(err);
        });

        io.emit('msg', data)
    });
});
