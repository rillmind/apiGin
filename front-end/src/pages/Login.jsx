import "../Global.css";
import "./Login.css";

function Login() {
  return (
    <div className="main">
      <div className="loginBox">
        <div className="login">
          <label htmlFor="login">Username ou Email</label>
          <input type="text" name="login" id="login" />
        </div>
        <div className="password">
          <label htmlFor="login">Senha</label>
          <input type="password" name="login" id="login" />
        </div>
        <div className="buttons">
          <a href="/">Entrar</a>
        </div>
      </div>
    </div>
  );
}

export default Login;
