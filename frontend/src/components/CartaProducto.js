import {React, useState} from 'react'
import { Card, Image, Header, Grid, Button, Icon, Modal, Form, Comment} from 'semantic-ui-react'
import "../css/CartaProducto.css"
import Comentario from './Comentario'
const axios = require('axios').default

function CartaProducto(props) {
    const [Unidades, setUnidades] = useState(1)
    const [comments, setComments] = useState([])
    const [open, setOpen] = useState(false)

    const AddCarrito = ()=>{
        var store = JSON.parse(localStorage.getItem('tienda'))
        var pedido = {
            Nombre:props.Nombre,
            Codigo:props.Codigo,
            Precio:parseFloat(props.Precio),
            Imagen:props.Imagen,
            Cantidad:Unidades
        }
        var datos = localStorage.getItem('carrito')
        if (datos == null|| datos == undefined ) {
            store = {
                Nombre: store.Nombre,
                Departamento: store.Departamento,
                Calificacion: store.Calificacion,
                Productos: [pedido]
            }
            localStorage.setItem('carrito',JSON.stringify([store]))

        }else{
            datos = JSON.parse(datos)
            var guardo = false
            datos.forEach(dato=>{
                if (dato.Nombre === store.Nombre && dato.Departamento === store.Departamento && dato.Calificacion === store.Calificacion) {
                    dato.Productos.forEach(producto => {
                        if (pedido.Codigo === producto.Codigo) {
                            producto.Cantidad += pedido.Cantidad
                            if(producto.Cantidad > parseInt(props.Cantidad)){
                                producto.Cantidad = parseInt(props.Cantidad)
                            }
                            guardo = true
                        }
                    })
                    if (!guardo) {
                        dato.Productos.push(pedido)
                        guardo = true
                    }
                }
            })
            if (!guardo) {
                store = {
                    Nombre: store.Nombre,
                    Departamento: store.Departamento,
                    Calificacion: store.Calificacion,
                    Productos: [pedido]
                }
                datos.push(store)
            }
            localStorage.setItem('carrito',JSON.stringify(datos))
        }    
    }
    if (props.Cantidad>0) {
        return (
            <>
            <Card onClick={()=>setOpen(!open)}>
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
                <Card.Meta><span className='date'>SKU: {props.Codigo}</span></Card.Meta>
                <Card.Description>{props.Descripcion}</Card.Description>
                </Card.Content>
                <Card.Content extra>
                <Button.Group floated="left" color='blue'>
                    <Button onClick={()=>{
                        if (Unidades>1) {
                            setUnidades(Unidades-1)
                        }
                    }}>-</Button>
                    <Button disabled>{Unidades}</Button>
                    <Button onClick={()=>{
                        if(Unidades<props.Cantidad){
                            setUnidades(Unidades+1)
                        }
                    }}>+</Button>
                </Button.Group>
                <Button floated="right" color="teal" onClick={AddCarrito}>
                    <Icon name='shop' />
                </Button>
                </Card.Content>
            </Card>
            <Modal
                open={open}
                onClose={() => setOpen(!open)}
            >
                <Modal.Content>
                <Comment.Group>
                        <Header as='h3' dividing>
                        Comentarios
                        </Header>
                        {comments.map((c,index)=>
                            <Comentario
                            Dpi={c.Comentario.Dpi}
                            Fecha={c.Comentario.Fecha}
                            Hora={c.Comentario.Hora}
                            Mensaje={c.Comentario.Mensaje}
                            SubComentarios={c.SubComentarios}
                            //Responder = {responder}
                            key = {index}
                            />
                        )}
                        <Form reply>
                        <Form.TextArea 
                            placeholder='Escribe un comentario...'
                            style={{minHeight:100, maxHeight:100}} 
                            onChange={(e)=>{
                                //setMessage(e.target.value)
                            }}
                        />
                            <Button content='Comentar' labelPosition='right' icon='edit' primary type='submit'/>
                        </Form>
                    </Comment.Group>
                </Modal.Content>
            </Modal>
            </>
            
        )
    }else{
        return(<></>)
    }
    
}

export default CartaProducto
