import axios from 'axios';


module.exports = {
    openWS:()=>{
        axios.get('/ws').then(res=>{
            console.log(res);
        }).catch(err=>{
            console.log(err);
        });
    }
}
