import {React, useEffect, useState} from 'react'
import {Icon, Header, Segment} from 'semantic-ui-react'
const axios = require('axios').default

function CargarArchivo(props) {
    const [Archivo, setArchivo] = useState(null)
    useEffect(() => {
        async function cargar(){
            if(Archivo!=null|undefined){
                setArchivo(null)
                let res = await axios.post(props.Ruta, Archivo);
                console.log(res)     
            }
        }
        cargar()
    })
    return (
        <div className="ui container">
            <Segment>
                <Header size="huge">
                    <Header.Content>{props.Title}</Header.Content>
                </Header>
            </Segment>
            <Segment placeholder>
                <Header icon>
                <Icon name='file code' />
                Elije un Archivo:
                </Header><br/>
                <input className="ui input" id="archivo" type="file" onChange={
                    (e)=>{
                        if (e.target.files[0]!=null){
                            let reader = new FileReader()
                            reader.readAsText(e.target.files[0], "UTF-8")
                            reader.onload=(a)=>{
                            //console.log(JSON.parse(a.target.result))
                            setArchivo(a.target.result)
                        }
                        }
                    }
                }/>
            </Segment>
        </div>
    )
}

export default CargarArchivo
