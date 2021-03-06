import { SHA256 } from 'crypto-js'
import {React, useState} from 'react'
import { useHistory } from 'react-router'
import { Button, Confirm, Form, Header, Icon, Input, Label, Segment } from 'semantic-ui-react'
import NavBar from './NavBar'
const axios = require('axios').default

function Registro() {
    let history = useHistory()
    const [dpi, setDpi] = useState('')
    const [password, setPassword] = useState('')
    const [confirmar, setConfirmar] = useState('')
    const [nombre, setNombre] = useState('')
    const [correo, setCorreo] = useState('')
    const [errDpi, setErrDpi] = useState('')
    const [errConfirmar, setErrConfirmar] = useState('')
    localStorage.clear()
    const [open, setOpen] = useState(false)
    let abrir = () => setOpen(true)
    let cerrar = () => setOpen(false)
    const registrar = ()=>{
        async function post(){
            let res = await axios.post('http://localhost:3000/Registro', {
                Dpi:parseInt(dpi),
                Nombre:nombre,
                Correo:correo,
                Password:SHA256(password).toString()
            })
            if (res.data.Tipo !== "Error"){
                localStorage.setItem("LOGED",'Cliente')
                localStorage.setItem("LOGUSER", dpi)
                history.push('/Home')
            }else{
                setErrDpi('*Usuario ya registrado')
            }
        }
        if (isNaN(parseInt(dpi))){
            setErrDpi('*DPI ingresado no válido.')
        }else{
            if (confirmar === password && confirmar!=="" && password!==""){
                post()
            }else{
                setErrConfirmar('*Las contraseñas no coinciden.')
            }
        }
        
    }

    return (
        <>
        <NavBar
        activo={[1]}
        />
        <div className="Content">
            <div className="ui segment mosaico container">
                <Segment>
                    <Header size="huge">
                        <Icon name='user plus'/>
                        <Header.Content>Registro</Header.Content>
                    </Header>
                </Segment>
                <Form onSubmit={registrar}>
                    <Form.Field>
                        <label>DPI:</label>
                        <Input required icon='address card' iconPosition='left' size="big" placeholder="DPI..." onChange={ (e)=>{
                                        setDpi(e.target.value);setErrDpi('')}}
                        />
                    </Form.Field>
                    <p style={{color:"red"}}>{errDpi}</p>
                    <Form.Field>
                        <label>Nombre:</label>
                        <Input required icon='pencil alternate' iconPosition='left' size="big" placeholder="Nombre..." onChange={ (e)=>{
                                        setNombre(e.target.value)}}/>
                    </Form.Field>
                    <Form.Field>
                        <label>Correo Electrónico:</label>
                        <Input required icon='mail outline' type='email' iconPosition='left' size="big" placeholder="Correo..." onChange={ (e)=>{
                                        setCorreo(e.target.value)}}/>
                    </Form.Field>
                    <Form.Field>
                        <label>Contraseña:</label>
                        <Input required type='password' icon='key' iconPosition='left' size="big" placeholder="Contraseña..." onChange={ (e)=>{
                                        setPassword(e.target.value);setErrConfirmar('')}}
                        />
                    </Form.Field>     
                    <Form.Field>
                        <label>Confirmar Contraseña:</label>
                        <Input required type='password' icon='key' iconPosition='left' size="big" placeholder="Confirmar Contraseña..." onChange={ (e)=>{
                                        setConfirmar(e.target.value);setErrConfirmar('')}}
                        />
                    </Form.Field>
                    <p style={{color:"red"}}>{errConfirmar}</p>
                    <center>
                        <Button type='submit' color="purple" size="big">Registrarme</Button>
                    </center>
                </Form>
            </div>
        </div>
        </>
    )
}

export default Registro
