import {React, useEffect, useState} from 'react'
import {Segment, Image, Loader} from 'semantic-ui-react'
import MosaicoTienda from './MosaicoTienda'
import '../css/Content.css'
const axios = require('axios').default

function StoreList() {
    const [Datos, setDatos] = useState([])
    const [req, setreq] = useState(false)
    var info = []
    useEffect(() => {
        async function obtener(){  
            if(!req){
                setreq(true)
                const data=await axios.get('http://localhost:3000/getTiendas')
                data.data.Datos.forEach(Dato => {
                    Dato.Departamentos.forEach(Departamento => {
                        Departamento.Tiendas.forEach(Tienda => {
                            info.push({
                                Nombre: Tienda.Nombre,
                                Descripcion: Tienda.Descripcion,
                                Contacto: Tienda.Contacto, 
                                Calificacion: Tienda.Calificacion,
                                Departamento: Departamento.Nombre,
                                Logo: Tienda.Logo
                            })
                        });
                    });
                });
                setDatos(info)              
            }
        }
        obtener()
    })
    if (Datos.length>0) {
        return (
            <div className="Content">
                <MosaicoTienda Datos={Datos}/>
            </div>
        )    
    }else{
        return(
            <div>
                <Segment>
                    <Loader active />
                    <Image src='https://react.semantic-ui.com/images/wireframe/short-paragraph.png' />
                </Segment>
            </div>
        )
    }
}

export default StoreList
