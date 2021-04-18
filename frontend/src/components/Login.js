import {React, useState} from 'react'
import { useHistory } from 'react-router'
import { Button, Confirm, Form, Header, Icon, Input, Segment, TransitionablePortal } from 'semantic-ui-react'
import NavBar from './NavBar'
import SHA256 from 'crypto-js/sha256';
const axios = require('axios').default

function Login() {
    localStorage.clear()
    const [dpi, setDpi] = useState("")
    const [password, setPassword] = useState("")
    const [errAbrir, setErrOpen] = useState(false)
    const [message, setMessage] = useState("")
    let history = useHistory()

    let errClose = () => setErrOpen(false)
    const ingresar = ()=>{
        async function obtener(){
            let res = await axios.post('http://localhost:3000/Login', {Dpi:parseInt(dpi),Password:SHA256(password).toString()})
            if (res.data.Tipo !== "Error"){
                localStorage.setItem("LOGED",res.data.Content)
                localStorage.setItem("LOGUSER", dpi)
                if (res.data.Content === "Admin"){
                    history.push("/Reporte")
                }else{
                    history.push("/Home")
                }
            }else{
                setMessage(res.data.Content)
                setErrOpen(true)
            }
        }
        obtener()
    }
    return (
        <>
        <NavBar
        activo={[0]}
        />
        <div className="Content">
            <div className="ui segment container">
                <Segment>
                    <Header size="huge">
                        <Icon name='user'/>
                        <Header.Content>Login</Header.Content>
                    </Header>
                </Segment>
                <center>
                    <Form>
                        <Input icon='address card' iconPosition='left' size="big" placeholder="DPI..." onChange={ (e)=>{
                                            setDpi(e.target.value)}}/>
                        <br/><br/>
                        <Input type='password' icon='key' iconPosition='left' size="big" placeholder="Contraseña..." onChange={ (e)=>{
                                            setPassword(e.target.value)}}/>
                        <br/><br/>
                        <Button type="submit" color="purple" size="big" onClick={ingresar}>Entrar</Button>
                    </Form>
                </center>
            </div>
        </div>
        <TransitionablePortal
            open={errAbrir}
            onClose={errClose}
        >   
            <Segment
            style={{ left: '40%', position: 'fixed', top: '50%', zIndex: 1000 }}
            >
            <Header>Error</Header>
            <p>{message}</p>
            </Segment>
        </TransitionablePortal>
        </>
    )
}

export default Login
