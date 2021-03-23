import {React, useState} from 'react'
import { Card, Image, Header, Grid, Button, Icon} from 'semantic-ui-react'
import "../css/CartaProducto.css"

function CartaProducto(props) {
    const [Carrito, setCarrito] = useState(1)
    return (
        <Card>
            <Card.Content extra>
                <Grid columns={2} relaxed='very' stackable>
                    <Grid.Column>
                        <center>
                            <Header sub>Precio:</Header>
                            <span className="Price">Q {props.Precio}</span>
                        </center>
                    </Grid.Column>
                    <Grid.Column>
                        <center>
                            <Header sub>Disponibilidad:</Header>
                            <span>{props.Cantidad}</span>
                        </center>
                    </Grid.Column>
                </Grid>
            </Card.Content>
            <Image src={props.Imagen} wrapped ui={false} />
            <Card.Content>
            <Card.Header>{props.Nombre}</Card.Header>
            <Card.Meta><span className='date'>{props.Codigo}</span></Card.Meta>
            <Card.Description>{props.Descripcion}</Card.Description>
            </Card.Content>
            <Card.Content extra>
                <Grid columns={2} relaxed='very' stackable>
                    <Grid.Column>
                        <center>
                            <Button.Group color='blue'>
                                <Button onClick={()=>{
                                    if (Carrito>1) {
                                        setCarrito(Carrito-1)
                                    }
                                }}>-</Button>
                                <Button disabled>{Carrito}</Button>
                                <Button onClick={()=>{
                                    if(Carrito<props.Cantidad){
                                        setCarrito(Carrito+1)
                                    }
                                }}>+</Button>
                            </Button.Group>
                        </center>
                    </Grid.Column>
                    <Grid.Column>
                        <center>
                            <Button secondary>
                                <Icon name='shop' />
                            </Button>
                        </center>
                    </Grid.Column>
                </Grid>
            </Card.Content>
        </Card>
    )
}

export default CartaProducto
