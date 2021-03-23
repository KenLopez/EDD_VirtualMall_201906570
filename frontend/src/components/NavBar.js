import {React, useState} from 'react'
import {Menu, Header, Image, Segment} from 'semantic-ui-react'
import {Link} from 'react-router-dom'
import '../css/NavBar.css'

const options=["Home", "Cargar Archivo", "Carrito de Compra", "Reportes"]
const url=["/Home", "/CargarArchivo", "/CarritoDeCompra", "/Reporte"]

function NavBar() {
    const [activo, setactivo] = useState(options[0])
    return (
        <>
        <Segment inverted color='purple' className="Header">
            <Header className="Title">
                <Image className="img" src='https://upload.wikimedia.org/wikipedia/commons/4/4a/Usac_logo.png' />
                Virtual Mall
            </Header>
            <Menu inverted className="Nav">
                {options.map((n,index)=>(
                    <Menu.Item as={Link} to={url[index]}
                        key={n}
                        name={n}
                        active={activo===n}
                        color={'purple'}
                        onClick={()=>setactivo(n)}
                    />
                ))}
            </Menu>
        </Segment>
        </>
    )
}

export default NavBar
