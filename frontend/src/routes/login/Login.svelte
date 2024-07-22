<script>
    import logo from '$lib/images/logo.svg';
</script>
<style>
    #container{
        width: 100vw;
        font-size: 2em;
        background-color: rgb(88, 88, 88);
        color: rgb(240,240,240);
        padding-left: 20px;
        flex: 1;
        box-sizing: border-box;
        padding-left: 38%;
        padding-top: 11%;
    }

    form{
        display: inline-block;
        width: 260px;
        height: 40%;
    }

    img{
        display: inline-block;
        width: 200px;
        height: 200px;
        text-align: center;
    }

    input[type="text"], input[type="password"] , input[type="submit"]{
        box-sizing: border-box;
        width: 220px;
        height: 36px;
        border-radius: 5px;
        padding-left: 10px;
        font-size: 16px;
        margin-top: 10px;
    }

    input[type="text"]:focus, input[type="password"]:focus{
        border: 2px solid black;
        outline: 0px;
    }

    input[type="submit"]{
        background-color: rgb(82, 192, 67);
        width: 160px;
        margin-left: 30px;
    }

    input[type="checkbox"]{
        width: 18px;
        height: 18px;
        margin: 2px;
        margin-left: 8px;
    }

    a{
        font-size: 20px;
        color: rgb(103, 20, 247)
    }
</style>

<div id="container">
    <button>Register</button><br>
    <form method="post">
        <input type="text" id="login" name="login" placeholder="Login"><br>
        <input type="password" id="password" name="password" placeholder="Password"><br>
        <input type="checkbox" id="if_remember" name="if_remember" value="RememberMe">
        <label for="if_remember" style="font-size: 50%">Remember me</label><br>
        <input type="submit" id="submit" name="submit" value="Submit"><br>
    </form>
    <img src={logo} alt="Logo isn't working"><br>
    <a href="">Help, I cannot log in</a><br>
    <a href="../yourservers">Skip</a>
</div>