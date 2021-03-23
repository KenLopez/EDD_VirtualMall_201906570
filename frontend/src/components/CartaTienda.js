import React from 'react'
import {Link} from 'react-router-dom'
import { Card, Image, Rating } from 'semantic-ui-react'

function CartaTienda(props) {
    return (
        <Card as={Link} to={`/Tienda/${props.Departamento}/${props.Nombre}/${props.Calificacion}`}>
            <Image src={props.Logo} wrapped ui={false} />
            <Card.Content>
            <Card.Meta><span className='date'>{props.Departamento}</span></Card.Meta>
            <Card.Header>{props.Nombre}</Card.Header>
            <Card.Meta><span className='date'>{props.Contacto}</span></Card.Meta>
            <Card.Description>{props.Descripcion}</Card.Description>
            </Card.Content>
            <Card.Content extra>
                <center>Calificación: <Rating icon='star' defaultRating={props.Calificacion} maxRating={5} disabled/></center>
            </Card.Content>
        </Card>
    )
}

export default CartaTienda
