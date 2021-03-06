import React, { useState } from 'react'
import {Link} from 'react-router-dom'
import { Card, Image, Rating } from 'semantic-ui-react'

function CartaTienda(props) {
    const [link, setLink] = useState(`/Tienda/${props.Departamento}/${props.Nombre}/${props.Calificacion}`)
    const [req, setReq] = useState(false)
    var desc = props.Descripcion.substring(0,100)
    if (desc.length == 100) {
        desc += '...'
    }
    if (!req){
        setReq(true)
        if (localStorage.getItem("LOGED") == 'Admin'){
            setLink(`/Inventarios/${props.Departamento}/${props.Nombre}/${props.Calificacion}`)
        }
    }
    return (
        <Card as={Link} to={link}>
            <Image src={props.Logo} wrapped ui={false} />
            <Card.Content>
            <Card.Meta><span className='date'>{props.Departamento}</span></Card.Meta>
            <Card.Header>{props.Nombre}</Card.Header>
            <Card.Meta><span className='date'>{props.Contacto}</span></Card.Meta>
            <Card.Description>{desc}</Card.Description>
            </Card.Content>
            <Card.Content extra>
                <center>Calificación: <Rating icon='star' defaultRating={props.Calificacion} maxRating={5} disabled/></center>
            </Card.Content>
        </Card>
    )
}

export default CartaTienda
