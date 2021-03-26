import React from 'react'
import {Button, Grid, Header, Image} from 'semantic-ui-react'

function ProductoPedido(props) {
    return (
        <>
            <Grid.Column>
                <center>
                    <Image src={props.Producto.Imagen} size='small'></Image>
                </center>
            </Grid.Column>
            <Grid.Column>
                <Header size='small'>
                    {props.Producto.Codigo} - {props.Producto.Nombre}    
                </Header>
            </Grid.Column>
            <Grid.Column>
                <Header size='small'>{props.Producto.Cantidad}</Header>
            </Grid.Column>
            <Grid.Column>
                <Header size='small'>{props.Producto.Precio}</Header>
            </Grid.Column>
            <Grid.Column>
                <Button negative onClick={()=>
                    props.borrarProducto(props.Producto.Codigo)
                }>X</Button>
            </Grid.Column>
        </>
    )
}

export default ProductoPedido
