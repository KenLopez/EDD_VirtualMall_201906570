import { SHA256 } from 'crypto-js'
import React, { useState } from 'react'
import { useHistory } from 'react-router'
import { Button, Confirm, Form, Header, Icon, Input, Segment, TransitionablePortal } from 'semantic-ui-react'
import NavBar from './NavBar'
const axios = require('axios').default

function Cuenta() {
    let history = useHistory()
    const [password2, setPassword2] = useState("")
    const [password, setPassword] = useState("")
    const [errAbrir, setErrOpen] = useState(false)
    const [message, setMessage] = useState("")
    const [open, setOpen] = useState(false)
    let abrir = () => setOpen(true)
    let cerrar = () => setOpen(false)

    function validar(){
        if(password === password2){
            setOpen(true)
        }else{
            setErrOpen(true)
            setMessage('Las contraseñas no coinciden.')
        }
    }

    let errClose = () => setErrOpen(false)
    const borrar = ()=>{
        async function obtener(){
            let res = await axios.delete('http://localhost:3000/EliminarCuenta',
            {
                data: {
                    Dpi:parseInt(localStorage.getItem('LOGUSER')),
                    Password:SHA256(password).toString()
                }
            })
            if (res.data.Tipo !== "Error"){
                history.push('/Login')
            }else{
                setOpen(false)
                setMessage("Credenciales incorrectas.")
                setErrOpen(true)
            }
        }
        obtener()
    }
    return (
        <>
        <NavBar
        activo={-1}
        />
        <div className="Content">
            <div className="ui segment container">
                <Segment>
                    <Header size="huge">
                        <Icon name='remove user'/>
                        <Header.Content>Eliminar Cuenta</Header.Content>
                    </Header>
                </Segment>
                <center>
                    <Form onSubmit={validar}>
                        <Input required type='password' icon='key' iconPosition='left' size="big" placeholder="Contraseña..." onChange={ (e)=>{
                                            setPassword(e.target.value)}}/>
                        <br/><br/>
                        <Input required type='password' icon='key' iconPosition='left' size="big" placeholder="Confirmar Contraseña..." onChange={ (e)=>{
                                            setPassword2(e.target.value)}}/>
                        <br/><br/>
                        <Button type="submit" color="red" size="big">Eliminar Mi Cuenta</Button>
                    </Form>
                </center>
            </div>
        </div>
        <Confirm
            cancelButton='Cancelar'
            confirmButton="Estoy Seguro"
            open={open}
            onCancel={cerrar}
            onConfirm={borrar}
            content="¿Seguro que deseas eliminar tu cuenta permanentemente?"
            size='mini'
        />
        <TransitionablePortal
        open={errAbrir}
        onClose={errClose}>   
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

export default Cuenta
