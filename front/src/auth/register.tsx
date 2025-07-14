import { useState } from 'react'
import { Link } from 'react-router-dom';

type Error = {
    message :string;
    isError: boolean;
}
export default function Register() {
    const [errorPwd, setErrorPwd] = useState<Error>({message:"", isError: false})
    const [errorEmail, setErrorEmail] = useState<Error>({message:"", isError: false})
    const [errorName, setErrorName] = useState<Error>({message:"", isError: false})
    const [errorPhone, setErrorPhone] = useState<Error>({message:"", isError: false})
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [fullname, setFullname] = useState('')
    const [phone, setPhone] = useState('test')

    const handleSave = function() {
        if (!password) {
            setErrorPwd({
                message: "Mot de passe est obligatoire",
                isError: true
            })
        } 
        if (!username) {
            setErrorEmail({
                message: "Email est obligatoire",
                isError: true
            })
        } 
        if (!phone) {
            setErrorPhone({
                message: "Téléphone est obligatoire",
                isError: true
            })
        } 
        if (!fullname) {
            setErrorName({
                message: "Nom et prénom est obligatoire",
                isError: true
            })
        } 
    }
    return (
        <div className="wrapper_register">
            <div className="form-group">
                <div className="error">{errorEmail.message}</div>
                <input type="text" className={(errorEmail?.isError ? "error" : "") + " form-control"}  id="username" placeholder="Email *" />
            </div>
            <div className="form-group">
                <div className="error">{errorPwd.message}</div>
                <input  type="password" className={(errorPwd?.isError ? "error" : "") + " form-control"} id="password"  placeholder="Mot de passe *" />
            </div>
            <div className="form-group">
                <div className="error">{errorName.message}</div>
                <input  type="text" className={(errorName?.isError ? "error" : "") + " form-control"} id="fullname"  placeholder="Nom complet *" />
            </div>
            <div className="form-group">
                <div className="error">{errorPhone.message}</div>
                <input  type="text" className={(errorPhone?.isError ? "error" : "") + " form-control"} id="phone"  placeholder="Numéro téléphone *" />
            </div>
            <div className="wrapper_btn_register">
                <button onClick={handleSave} className="btn btn-primary">S'inscrire</button>
            </div>
            <div className="wrapper_link">
                <Link to="/login">Vous avez déjà un compte ?</Link>
            </div>
        </div>
    )
}