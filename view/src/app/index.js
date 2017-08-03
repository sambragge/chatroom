import React, {Component} from 'react';
import Socket from '../socket';
import backend from '../backend';

export default class App extends Component {

    constructor(props){
        super(props);
        this.state = {
            message:"",
            messages:[],
            users:0,
            connected:false,
        }
    }

    componentDidMount(){
        const self = this;
        this.socket = new Socket()

        this.socket.emit("client handshake")
        this.socket.on("server handshake", ()=>{
            console.log("recieved the server handshake")
        })
    }

    render(){
        return(
            <div className="app container">

                <div className="message-window row u-full-width"></div>

                <form className="inputs row">
                    <input className="message-input ten columns" placeholder="Enter message here..." type="text"/>
                    <input className="send-button two columns" type="submit" value="send"/>
                </form>

            </div>
        )
    }
}
