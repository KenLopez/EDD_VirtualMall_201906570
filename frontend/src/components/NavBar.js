import {React, useState} from 'react'
import {Menu, Header, Image, Segment, Button, Confirm} from 'semantic-ui-react'
import {Link, useHistory} from 'react-router-dom'
import '../css/NavBar.css'

//const url=["/Home", "/CargarArchivo", "/CarritoCompra", "/Reporte"]

function NavBar(props) {
    let history = useHistory()
    const [color, setColor] = useState('purple')
    const [req, setReq] = useState(false)
    const [title, setTitle] = useState('Virtual Mall')
    const [options, setOptions] = useState([])
    const [url, setUrl] = useState([])
    const [activo, setActivo] = useState(null)
    const [logout, setLogout] = useState(false)

    const out = ()=>{
        history.push('/Login')
    }
    if (!req){
        var option = []
        var link = []
        if (localStorage.getItem("LOGED") == 'Admin'){
            setColor('teal')
            setTitle('Virtual Mall Administrador')
            option=["Reportes", "Inventarios", "Cargar Archivos"]
            link = ["/Reporte", "/Inventarios", "/CargarArchivo", "/Login"]
            setLogout(true)
        }else if (localStorage.getItem("LOGED") == 'Cliente'){
            option=["Home", "Carrito de Compra","Cuenta"]
            link=["/Home", "/CarritoCompra"]
            setLogout(true)
        }else{
            option=["Iniciar Sesión", "Registro"]
            link=["/Login", "/Registro"]
        }
        setActivo(option[props.activo])
        setOptions(option)
        setUrl(link)
        setReq(true)
    }
    
    return (
        <>
        <Segment inverted color={color} className="Header">
            <Header className="Title">
                <Image className="img" src='https://upload.wikimedia.org/wikipedia/commons/4/4a/Usac_logo.png' />
                {title}
            </Header>
            <Menu inverted className="Nav">
                {options.map((n,index)=>(
                    <Menu.Item as={Link} to={url[index]}
                        key={n}
                        name={n}
                        active={activo===n}
                        color={color}
                        onClick={()=>setActivo(n)}
                    />
                ))}
                <Logout
                yes={logout}
                out={out}
                />
            </Menu>
        </Segment>
        </>
    )
}

function Logout(props){
    const [open, setOpen] = useState(false)
    let abrir = () => setOpen(true)
    let cerrar = () => setOpen(false)
    if (props.yes){
        return(
            <>
                <Menu.Item as={Button}
                    position='right'
                    key={"Logout"}
                    name={"Logout"}
                    color='red'
                    active={true}
                    onClick={abrir}
                />
                <Confirm
                    cancelButton='Cancelar'
                    confirmButton="Salir de Virtual Mall"
                    open={open}
                    onCancel={cerrar}
                    onConfirm={props.out}
                    content="¿Seguro que desea salir?"
                    size='mini'
                />
            </>
        )
    }else{
        return <></>
    }
}

export default NavBar
