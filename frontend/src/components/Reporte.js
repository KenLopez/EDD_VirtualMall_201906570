import {React, useState} from 'react'
import {Header, Segment, Icon, Container, Grid, Input, Button, Image, Dropdown} from 'semantic-ui-react'
import '../css/Content.css'
const axios = require('axios').default

function Reporte() {
    const [imagen, setImagen] = useState('')
    const [title, setTitle] = useState('')
    const [year, setYear] = useState('')
    const [month, setMonth] = useState('0')
    const [cat, setCat] = useState('')
    const [day, setDay] = useState('')
    function getMonth(m){
        switch (m) {
            case 'ENERO':
                return '1'
            case 'FEBRERO':
                return '2'
            case 'MARZO':
                return '3'
            case 'ABRIL':
                return '4'
            case 'MAYO':
                return '5'
            case 'JUNIO':
                return '6'
            case 'JULIO':
                return '7'
            case 'AGOSTO':
                return '8'
            case 'SEPTIEMBRE':
                return '9'
            case 'OCTUBRE':
                return '10'
            case 'NOVIEMBRE':
                return '11'
            case 'DICIEMBRE':
                return '12'
            default:
                return '0';
        }
    }
    const monthOptions = [
        {key: 1, text: 'ENERO', value: 1},
        {key: 2, text: 'FEBRERO', value: 2},
        {key: 3, text: 'MARZO', value: 3},
        {key: 4, text: 'ABRIL', value: 4},
        {key: 5, text: 'MAYO', value: 5},
        {key: 6, text: 'JUNIO', value: 6},
        {key: 7, text: 'JULIO', value: 7},
        {key: 8, text: 'AGOSTO', value: 8},
        {key: 9, text: 'SEPTIEMBRE', value: 9},
        {key: 10, text: 'OCTUBRE', value: 10},
        {key: 11, text: 'NOVIEMBRE', value: 11},
        {key: 12, text: 'DICIEMBRE', value: 12},
    ]
    const arbolA = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetArbolAnio')
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Árbol Años')
            }
        }
        obtener()
    }
    const arbolM = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetArbolMeses/'+year)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Árbol Meses')
            }
        }
        obtener()
    }
    const matriz = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetMatriz/'+year+'/'+month)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Matriz Pedidos')
            }
        }
        obtener()
    }
    const cola = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetPedidosDia/'+year+'/'+month+'/'+cat+'/'+day)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Cola de Pedidos')
            }
        }
        obtener()
    }
    return (
        <div className="Content">
            <div className="ui segment mosaico container">
                <Segment>
                    <Header size="huge">
                        <Icon name='line graph'/>
                        <Header.Content>Reportes</Header.Content>
                    </Header>
                </Segment>
                <Container fluid>
                    <Grid>
                        <Grid.Row columns={1}>
                            <Grid.Column>
                                <Button color='teal' fluid onClick={arbolA}>Obtener Árbol Años</Button>
                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row columns={4}>
                            <Grid.Column>
                               <Input fluid placeholder='Año...' onChange={ (e)=>{
                                        setYear(e.target.value)
                                    }}/>
                            </Grid.Column>
                            <Grid.Column>
                               <Dropdown placeholder='Mes...' selection fluid options={monthOptions} onChange={ (e)=>{
                                        setMonth(getMonth(e.target.innerText))
                                    }}/> 
                            </Grid.Column>
                            <Grid.Column>
                               <Input fluid placeholder='Categoria...'onChange={ (e)=>{
                                        setCat(e.target.value)
                                    }}/>
                            </Grid.Column>
                            <Grid.Column>
                               <Input fluid placeholder='Día...'onChange={ (e)=>{
                                        setDay(e.target.value)
                                    }}/>
                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row columns={3}>
                            <Grid.Column>
                                <center>
                                    <Button color='teal' onClick={arbolM}>Obtener Árbol<br/>de Meses</Button>
                                </center>
                            </Grid.Column>
                            <Grid.Column>
                                <center>
                                    <Button color='teal' onClick={matriz}>Obtener Matriz<br/>de Pedidos</Button>
                                </center>
                            </Grid.Column>
                            <Grid.Column >
                                <center>
                                    <Button color='teal' onClick={cola}>Obtener Cola<br/>de Pedidos</Button>
                                </center>
                            </Grid.Column>
                        </Grid.Row>
                    </Grid>
                </Container>
                <Segment>
                    <Header size="huge">
                        <Header.Content>{title}</Header.Content>
                    </Header>
                </Segment>
                <Image fluid src={imagen}/>
            </div>
        </div>
        
    )
}
export default Reporte
