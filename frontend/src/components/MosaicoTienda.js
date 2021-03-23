import {React} from 'react'
import { Segment, Header,Icon } from 'semantic-ui-react'
import CartaTienda from './CartaTienda'

function MosaicoTienda(props) {
    return (
        <div className="ui segment mosaico container">
            <Segment>
                <Header size="huge">
                    <Icon name='dollar sign'/>
                    <Header.Content>Tiendas</Header.Content>
                </Header>
            </Segment>
            <div className="ui four column link cards row">
                {props.Datos.map((c,index)=>
                    <CartaTienda
                        Nombre={props.Datos[index].Nombre}
                        Descripcion={props.Datos[index].Descripcion}
                        Contacto={props.Datos[index].Contacto}
                        Calificacion={props.Datos[index].Calificacion}
                        Logo={props.Datos[index].Logo}
                        Departamento={props.Datos[index].Departamento}
                        key={index}
                    />
                )
                }
            </div>
        </div>
    )
}

export default MosaicoTienda
