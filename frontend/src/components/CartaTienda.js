import React from 'react'
import { Card, Image, Rating } from 'semantic-ui-react'

function CartaTienda(props) {
    return (
        <Card>
            <Image src={props.Logo} wrapped ui={false} />
            <Card.Content>
            <Card.Meta><span className='date'>{props.Departamento}</span></Card.Meta>
            <Card.Header>{props.Nombre}</Card.Header>
            <Card.Meta><span className='date'>{props.Contacto}</span></Card.Meta>
            <Card.Description>{props.Descripcion}</Card.Description>
            </Card.Content>
            <Card.Content extra>
                <Rating icon='star' defaultRating={props.Calificacion} maxRating={5} disabled/>
            </Card.Content>
        </Card>
    )
}

export default CartaTienda
