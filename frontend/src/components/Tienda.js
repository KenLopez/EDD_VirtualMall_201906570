import {React, useEffect, useState} from 'react'
import { useHistory, useParams } from 'react-router'
import {Segment, Header, Icon, Loader, Image, Grid, Rating, Container, Message, Comment, Form, Button} from 'semantic-ui-react'
import CartaProducto from './CartaProducto'
import '../css/Content.css'
import NavBar from './NavBar'
const axios = require('axios').default

function Tienda() {
    const history = useHistory()
    if (localStorage.getItem("LOGED") == null){
        history.push("/Login")
    }else if (localStorage.getItem("LOGED")=="Admin"){
        history.push("/Reporte")
    }
    let {Departamento, Nombre, Calificacion} = useParams()
    const [Datos, setDatos] = useState([])
    const [Descripcion, setDescripcion] = useState('')
    const [Contacto, setContacto] = useState('')
    const [data, setData] = useState('')
    const [req, setreq] = useState(false)
    var info = []
    var tienda = {
        Nombre:Nombre,
        Departamento:Departamento,
        Calificacion:Calificacion,
    }
    localStorage.setItem('tienda', JSON.stringify(tienda))
    useEffect(() => {
        async function obtener(){
            if(!req){
                setreq(true)
                let res = await axios.post('http://localhost:3000/GetInventario', { Departamento: Departamento, Nombre: Nombre, Calificacion: parseInt(Calificacion,10) })
                info.push(res.data)
                //info.push(res2.data)
                //console.log(info)
                setDatos(info[0].Productos) 
                setDescripcion(info[0].Descripcion)
                setContacto(info[0].Contacto)
                if (info.length>0) {
                    let res2 = await axios.post('http://localhost:3000/GetArbolInventario', { Departamento: Departamento, Nombre: Nombre, Calificacion: parseInt(Calificacion,10) })
                    setData("data:image/png;base64,"+res2.data)   
                }
                //setDot(info[1])             
            }
        }
        obtener()
    })
    if (Datos.length>0) {
        return (
            <>
            <NavBar/>
            <div className="Content">
                <div className="ui segment mosaico container">
                    <Segment>
                            <Grid columns={2} relaxed='very' stackable>
                                <Grid.Column>
                                    <Header size="huge">
                                        <Icon name='shopping bag'/>
                                        <Header.Content>{Nombre}</Header.Content>
                                    </Header>
                                </Grid.Column>
                                <Grid.Column>
                                    <Header size="huge" textAlign="right">
                                        <Header.Content>{Departamento}</Header.Content>
                                    </Header>
                                </Grid.Column>
                            </Grid>
                            <Header size="huge">
                                <Rating icon='star' defaultRating={Calificacion} maxRating={5} disabled/>
                            </Header>
                            <Header size="medium">
                                <Header.Content>{Contacto}</Header.Content>
                            </Header>
                            <br/>
                            <Container textAlign="center" fluid>
                                <Header size="medium">{Descripcion}</Header>
                            </Container>
                    </Segment>
                    <div className="ui three column link cards row">
                        {Datos.map((c,index)=>
                                <CartaProducto
                                    Nombre={Datos[index].Nombre}
                                    Descripcion={Datos[index].Descripcion}
                                    Codigo={Datos[index].Codigo}
                                    Precio={Datos[index].Precio}
                                    Imagen={Datos[index].Imagen}
                                    Cantidad={Datos[index].Cantidad}
                                    key={Datos[index].Codigo}
                                />
                            )
                        }
                    </div>
                    <Comment.Group>
                        <Header as='h3' dividing>
                        Comentarios
                        </Header>

                        <Comment>
                        <Comment.Avatar src='https://react.semantic-ui.com/images/avatar/small/matt.jpg' />
                        <Comment.Content>
                            <Comment.Author as='a'>Matt</Comment.Author>
                            <Comment.Metadata>
                            <div>Today at 5:42PM</div>
                            </Comment.Metadata>
                            <Comment.Text>How artistic!</Comment.Text>
                            <Comment.Actions>
                            <Comment.Action>Reply</Comment.Action>
                            </Comment.Actions>
                        </Comment.Content>
                        </Comment>

                        <Comment>
                        <Comment.Avatar src='https://react.semantic-ui.com/images/avatar/small/elliot.jpg' />
                        <Comment.Content>
                            <Comment.Author as='a'>Elliot Fu</Comment.Author>
                            <Comment.Metadata>
                            <div>Yesterday at 12:30AM</div>
                            </Comment.Metadata>
                            <Comment.Text>
                            <p>This has been very useful for my research. Thanks as well!</p>
                            </Comment.Text>
                            <Comment.Actions>
                            <Comment.Action>Reply</Comment.Action>
                            </Comment.Actions>
                        </Comment.Content>
                        <Comment.Group>
                            <Comment>
                            <Comment.Avatar src='https://react.semantic-ui.com/images/avatar/small/jenny.jpg' />
                            <Comment.Content>
                                <Comment.Author as='a'>Jenny Hess</Comment.Author>
                                <Comment.Metadata>
                                <div>Just now</div>
                                </Comment.Metadata>
                                <Comment.Text>Elliot you are always so right :)</Comment.Text>
                                <Comment.Actions>
                                <Comment.Action>Reply</Comment.Action>
                                </Comment.Actions>
                            </Comment.Content>
                            </Comment>
                        </Comment.Group>
                        </Comment>

                        <Comment>
                        <Comment.Avatar src='https://react.semantic-ui.com/images/avatar/small/joe.jpg' />
                        <Comment.Content>
                            <Comment.Author as='a'>Joe Henderson</Comment.Author>
                            <Comment.Metadata>
                            <div>5 days ago</div>
                            </Comment.Metadata>
                            <Comment.Text>Dude, this is awesome. Thanks so much</Comment.Text>
                            <Comment.Actions>
                            <Comment.Action>Reply</Comment.Action>
                            </Comment.Actions>
                        </Comment.Content>
                        </Comment>

                        <Form reply>
                        <Form.TextArea />
                        <Button content='Add Reply' labelPosition='left' icon='edit' primary />
                        </Form>
                    </Comment.Group>
                </div>
            </div>
            </>
        )
    }else{
        if(!req){
            return(
                <>
                <NavBar/>
                <div>
                    <Segment>
                        <Loader active />
                        <Image src='https://react.semantic-ui.com/images/wireframe/short-paragraph.png' />
                    </Segment>
                </div>
                </>
            )
        }else{
            return(
                <>
                <NavBar/>
                <div className="Content">
                    <div className="ui segment mosaico container">
                        <Segment>
                                <Grid columns={2} relaxed='very' stackable>
                                    <Grid.Column>
                                        <Header size="huge">
                                            <Icon name='shopping bag'/>
                                            <Header.Content>{Nombre}</Header.Content>
                                        </Header>
                                    </Grid.Column>
                                    <Grid.Column>
                                        <Header size="huge" textAlign="right">
                                            <Header.Content>{Departamento}</Header.Content>
                                        </Header>
                                    </Grid.Column>
                                </Grid>
                                <Header size="huge">
                                    <Rating icon='star' defaultRating={Calificacion} maxRating={5} disabled/>
                                </Header>
                                <Header size="medium">
                                    <Header.Content>{Contacto}</Header.Content>
                                </Header>
                                <br/>
                                <Container textAlign="center" className="fluid">
                                    <Header size="medium">{Descripcion}</Header>
                                </Container>
                        </Segment>
                        <Message>
                            <Message.Header>Esta tienda no posee inventario</Message.Header>
                            {
                                localStorage.getItem("LOGED")==="Admin"?(
                                    <p>Puedes cargar inventarios en la sección de Cargar Archivo.</p>
                                ):(
                                    <></>
                                )
                            }
                            
                        </Message>
                    </div>
                </div>
                </>
            )
        }
    }
}

export default Tienda
