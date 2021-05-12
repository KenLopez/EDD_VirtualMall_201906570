import React, { useState } from 'react'
import { Button, Comment, Form } from 'semantic-ui-react'

function Comentario(props) {
    const [resp, setResp] = useState(false)
    const [message, setMessage] = useState('')
    const responder = (hijo)=>{
        let comment = {
            Comentario:{
                Dpi:props.Dpi,
                Mensaje:props.Mensaje,
                Fecha:props.Fecha,
                Hora:props.Hora
            },
            Sub:hijo
        }
        Array.from(document.querySelectorAll("textarea")).forEach(
            item => (item.value = "")
        );
        props.Responder(comment)
    }
    return (
        <Comment>
        <Comment.Content>
            <Comment.Author as='a'>{props.Dpi}</Comment.Author>
            <Comment.Metadata>
            <div>{props.Fecha} {props.Hora}</div>
            </Comment.Metadata>
            <Comment.Text>
            <p>{props.Mensaje}</p>
            </Comment.Text>
            <Comment.Actions>
            <Comment.Action as='a' onClick={()=>{setResp(!resp)}}>Reply</Comment.Action>
            </Comment.Actions>
        </Comment.Content>
        {resp?
        <Form reply>
        <Form.TextArea 
            placeholder='Escribe un comentario...'
            style={{minHeight:50, maxHeight:50}} 
            onChange={(e)=>{setMessage(e.target.value)}}
        />
        <Button content='Comentar' labelPosition='right' icon='edit' primary onClick={()=>{
            if (message!=='') {
                let m = message
                setMessage(m)
                setResp(!resp)
                responder(
                    {
                        Comentario:{
                            Dpi:parseInt(localStorage.getItem("LOGUSER")),
                            Mensaje:m
                        },
                        Sub:null
                    }
                )
            }
        }}/>
        </Form>
        :
        <></>}
        {((props.SubComentarios != null) && (props.SubComentarios.length>0))?
        <Comment.Group>
            {props.SubComentarios.map((c,index)=>
                <Comentario
                Dpi={c.Comentario.Dpi}
                Fecha={c.Comentario.Fecha}
                Hora={c.Comentario.Hora}
                Mensaje={c.Comentario.Mensaje}
                SubComentarios={c.SubComentarios}
                Responder={responder}
                key={index}
                />
            )}
        </Comment.Group>
        :
        <></>
        }
        </Comment>
    )
}

export default Comentario
