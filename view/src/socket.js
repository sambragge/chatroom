import {EventEmitter} from 'events';

export default class Socket {

    constructor(ws = new WebSocket('ws://localhost:8000/ws'), ee = new EventEmitter()){
        console.log("constructing Socket")

        this.ws = ws;
        this.ee = ee;


        this.ws.onmessage = this.message.bind(this);
        this.ws.onopen = this.open.bind(this);
        this.ws.onclose = this.close.bind(this);
    }

    on(name, func){
        this.ee.on(name, func);
    }

    off(name, func){
        this.ee.removeListener(name, func);
    }

    emit(name, data){
        const message = JSON.stringify({name, data});
        this.ws.send(message);
    }


    message(e){
        try{
            const message = JSON.parse(e.data);
            this.ee.emit(message.name, message.data);

        }
        catch(err){
            this.ee.emit('error', err);
        }

    }

    open(){
        this.ee.emit('connected');
    }

    close(){
        this.ee.emit('disconnect');
    }
}
