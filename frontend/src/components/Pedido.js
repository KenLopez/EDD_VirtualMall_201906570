import {React,useEffect, useState} from 'react'
import {Segment, Header, Grid, Button} from 'semantic-ui-react'
import ProductoPedido from './ProductoPedido'

function Pedido(props) {
    const [suma, setsuma] = useState(0)

    var borrarProducto = (dataHijo)=>{
        for (let i = 0; i < props.data.Productos.length; i++) {
            if(props.data.Productos[i].Codigo == dataHijo){
                var productosN = props.data.Productos
                productosN.splice(i,1)
                var dataN = {
                    Nombre: props.data.Nombre,
                    Departamento: props.data.Departamento,
                    Calificacion: props.data.Calificacion,
                    Productos: productosN
                }
                props.borrarProducto(dataN)
            }
        }
    }

    useEffect(() => {
        function sumar(){
            let total = 0
            props.data.Productos.forEach(producto => {
                total += producto.Precio * producto.Cantidad
            });
            setsuma(total)
        }
        sumar()
    })
    return (
        <Segment>
            <Grid divided='vertically'>
                <Grid.Row columns={2}>
                    <Grid.Column>
                        <Header size='medium'>{props.data.Nombre} (Calif.{props.data.Calificacion})</Header>
                    </Grid.Column>
                    <Grid.Column>
                        <Header textAlign='right' size='medium'>{props.data.Departamento}</Header>
                    </Grid.Column>
                </Grid.Row>
                <Grid.Row textAlign='center' columns={5}>
                    <Grid.Column>
                        <Header size='medium'>Imagen</Header>
                    </Grid.Column>
                    <Grid.Column>
                        <Header size='medium'>Nombre</Header>
                    </Grid.Column>
                    <Grid.Column>
                        <Header size='medium'>Cantidad</Header>
                    </Grid.Column>
                    <Grid.Column>
                        <Header size='medium'>Costo Unitario (Q)</Header>
                    </Grid.Column>
                    <Grid.Column></Grid.Column>
                    {props.data.Productos.map((c,index)=>
                        <ProductoPedido
                            Producto={c}
                            borrarProducto={borrarProducto}
                            key={index}
                        />
                    )}
                </Grid.Row>
                <Grid.Row columns={5}>
                    <Grid.Column>
                        <Header size='medium' textAlign='center'>TOTAL</Header>
                    </Grid.Column>
                    <Grid.Column></Grid.Column>
                    <Grid.Column></Grid.Column>
                    <Grid.Column>
                        <Header size='medium' textAlign='center'>Q {suma}</Header>
                    </Grid.Column>
                    <Grid.Column></Grid.Column>
                </Grid.Row>
                <Grid.Row columns={1}>
                    <Grid.Column>
                        <Button onClick={()=>{
                            props.confirmar(props.data)
                        }} floated='right' positive>
                            Confirmar Pedido
                        </Button>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Segment>
    )
}

export default Pedido
