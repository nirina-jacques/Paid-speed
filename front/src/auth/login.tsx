import { Link } from 'react-router-dom'

export default function Login() {
    return (
        <div className="login_wrapper">
            <div className="demo">
                <Link to="/">Ariary-To</Link>
            </div>
            <div className="wrapper_fb">
                <span className="color_wh">
                    Se connecter avec Facebook
                </span>
            </div>
            <div className="content_body">
                <div className="form-group">
                    <input type="text" className="form-control"  id="username" placeholder="Email" />
                </div>
                <div className="form-group">
                    <input  type="password" className="form-control" id="password"  placeholder="Mot de passe" />
                </div>
                <div className="forgot-pwd text-center">
                    <Link to="/">Mot de passe oublié ?</Link>
                </div>
                <div className="wrapper_btn_log">
                    <button  className="btn btn-primary">Se connecter</button>
                </div>
            </div>
            <div className="register-in-login">
                <div className="text-center">Pas encore de compte Ariary-To ?</div>
                <Link to="/register">Crée un compte</Link>
            </div>
    </div>
    )
}