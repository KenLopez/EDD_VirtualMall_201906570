import {React, useState} from 'react'
import {Menu} from 'semantic-ui-react'
import {Link} from 'react-router-dom'

const options=["Home", "Cargar Tiendas", "Cargar Productos", 
"Cargar Pedidos", "Reportes"]
const url=["/Home", "/CargarTienda", "/CargarProducto", "/CargarPedido", "/Reporte"]

function NavBar() {
    const [activo, setactivo] = useState(options[0])
    return (
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
    )
}

export default NavBar
