import {React, useState, useEffect} from 'react'
import {Segment, Header, Icon, Message, Button} from 'semantic-ui-react'
import Pedido from './Pedido'
import '../css/Content.css'
//const axios = require('axios').default

function CarritoCompra() {
    const [carritos, setcarritos] = useState([])
    const [req, setreq] = useState(false)

    var confirmar = (dataHijo)=>{
        for (let index = 0; index < carritos.length; index++) {
            if(carritos[index].Nombre === dataHijo.Nombre && carritos[index].Departamento === dataHijo.Departamento && carritos[index].Calificacion == dataHijo.Calificacion){
                var carritosN = carritos
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
        )
    }else{
        return(
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
                    <Button positive floated='right'>Confirmar Todos</Button>
                </div>
            </div>
        )
    }

}

export default CarritoCompra
