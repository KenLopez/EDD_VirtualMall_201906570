import {React, useState, useEffect} from 'react'
import {Segment, Header, Icon, Message, Button} from 'semantic-ui-react'
import Pedido from './Pedido'
import '../css/Content.css'
import NavBar from './NavBar'
import { useHistory } from 'react-router'
const axios = require('axios').default

function CarritoCompra() {
    const history = useHistory()
    if (localStorage.getItem("LOGED") == null){
        history.push("/Login")
    }else if (localStorage.getItem("LOGED")=="Admin"){
        history.push("/Reporte")
    }
    const [carritos, setcarritos] = useState([])
    const [req, setreq] = useState(false)
    
    var enviarPedido = (pedido)=>{
        async function enviar(){
            let res = await axios.post("http://localhost:3000/CargarPedidos", pedido);
            console.log(res)    
        }
        enviar()
    }

    var enviarTodos = ()=>{
        carritos.forEach(carrito => {
            var today = new Date()
            var archivo = {Pedidos:[{
                Fecha: String(today.getDate()).padStart(2, '0')+"-"+String(today.getMonth() + 1).padStart(2, '0')+"-"+today.getFullYear(),
                Tienda: carrito.Nombre,
                Departamento: carrito.Departamento,
                Calificacion: parseInt(carrito.Calificacion),
                Productos: [],
                Cliente: parseInt(localStorage.getItem('LOGUSER'))
            }]}
            carrito.Productos.forEach(producto => {
                for (let i = 0; i < producto.Cantidad; i++) {
                    archivo.Pedidos[0].Productos.push({Codigo: producto.Codigo})   
                }
            })
            enviarPedido(archivo)
        })
        setcarritos([])
        localStorage.setItem('carrito',"[]")
        setreq(false)
    }

    var confirmar = (dataHijo)=>{
        for (let index = 0; index < carritos.length; index++) {
            if(carritos[index].Nombre === dataHijo.Nombre && carritos[index].Departamento === dataHijo.Departamento && carritos[index].Calificacion == dataHijo.Calificacion){
                var carritosN = carritos
                var today = new Date()
                var archivo = {Pedidos:[{
                    Fecha: String(today.getDate()).padStart(2, '0')+"-"+String(today.getMonth() + 1).padStart(2, '0')+"-"+today.getFullYear(),
                    Tienda: carritosN[index].Nombre,
                    Departamento: carritosN[index].Departamento,
                    Calificacion: parseInt(carritosN[index].Calificacion),
                    Productos: [],
                    Cliente: parseInt(localStorage.getItem('LOGUSER'))
                }]}
                carritosN[index].Productos.forEach(producto => {
                    for (let i = 0; i < producto.Cantidad; i++) {
                        archivo.Pedidos[0].Productos.push({Codigo: producto.Codigo})   
                    }
                })
                enviarPedido(archivo)
                carritosN.splice(index,1)
                setcarritos(carritosN)
                localStorage.setItem('carrito', JSON.stringify(carritosN))
                setreq(false)
                break
            }
        }
    }

    var borrarProducto = (dataHijo)=>{
        for (let index = 0; index < carritos.length; index++) {
            if(carritos[index].Nombre === dataHijo.Nombre && carritos[index].Departamento === dataHijo.Departamento && carritos[index].Calificacion == dataHijo.Calificacion){
                var carritosN = carritos
                if (dataHijo.Productos.length>0){
                    carritosN[index].Productos = dataHijo.Productos
                }else{
                    carritosN.splice(index,1)
                }
                setcarritos(carritosN)
                localStorage.setItem('carrito', JSON.stringify(carritosN))
                setreq(false)
                break
            }
        }
    }

    useEffect(() => {
        function obtener(){
            if (!req) {
                setreq(true)
                var data = []
                data = JSON.parse(localStorage.getItem('carrito'))
                if (data != null) {
                    setcarritos(data)
                }
            }
        }
        obtener()
    })
    
    if (carritos.length === 0 && req) {
        return (
            <>
            <NavBar
            activo={1}
            />
            <div className="Content">
                <div className="ui segment mosaico container">
                    <Segment>
                        <Header size="huge">
                            <Icon name='shop'/>
                            <Header.Content>Carrito de Compra</Header.Content>
                        </Header>
                    </Segment>
                    <Message>
                        <Message.Header>Tu carrito de compras está vacío</Message.Header>
                        <p>Consigue artículos de las tiendas de Home.</p>
                    </Message>
                </div>
            </div>
            </>
        )
    }else{
        return(
            <>
            <NavBar
            activo={1}
            />
            <div className="Content">
                <div className="ui segment mosaico container">
                    <Segment>
                        <Header size="huge">
                            <Icon name='shop'/>
                            <Header.Content>Carrito de Compra</Header.Content>
                        </Header>
                    </Segment>
                    {carritos.map((c,index)=>
                        <Pedido
                            data={carritos[index]}
                            confirmar={confirmar}
                            borrarProducto={borrarProducto}
                            key={index}
                        />
                    )}
                    <Button onClick={()=>{
                        setcarritos([])
                        localStorage.clear()
                    }}>Limpiar</Button>
                    <Button positive floated='right' onClick={enviarTodos}>Confirmar Todos</Button>
                </div>
            </div>
            </>
        )
    }

}

export default CarritoCompra
