const cors = require("cors");
const express = require('express');
const app = express();
app.use(cors());

var mysql = require('mysql2')
var conn = mysql.createConnection({ // createConnection 메서드로 객체화
    port : '3306',
    host : '13.125.236.0',  
    user : 'glassKT',
    password : '1234',
    database : 'glasskt'
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
        console.log(data)
        
        socket.join(data.roomnum);
        io.to(data.roomnum).emit("msg", data.name + " is join " + data.roomnum)
    });

    socket.on('leaveroom', (data) => {
        console.log(data)

        socket.leave(data.roomnum);
        io.to(data.roomnum).emit("msg", data.name + " is leave " + data.roomnum)
    });


    socket.on('msg', (data) => {        // data { roomnum, content, id, name, createat }
        console.log(data);

        conn.query(sql,[data.num_room, data.content, data.user_id, data.createat, data.name], function(err, rows, fields){
            console.log(err);
        });

        socket.emit('msg', data)
    });
});
