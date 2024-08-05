<script>
    import logo from '$lib/images/logo.svg';
    import { Modal } from '@svelteuidev/core';
    let opened = false;
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
        padding-top: 10%;
    }
 
    #formContainer{
        width: 270px;
        display: inline-block;
        margin-right: 30px;
    }

    form{
        display: inline-block;
        width: 260px;
        height: 40%;
    }

    img{
        display: inline-block;
        width: 300px;
        height: 300px;
        text-align: center;
    }

    input[type="text"], input[type="password"] , input[type="submit"]{
        box-sizing: border-box;
        width: 100%;
        height: 40px;
        border-radius: 5px;
        padding-left: 10px;
        font-size: 16px;
        margin-top: 10px;
    }

    input[type="text"]:focus, input[type="password"]:focus{
        border: 2px solid black;
        outline: 0px;
        width: 100%;
    }

    input[type="submit"]{
        background-color: rgb(82, 192, 67);
        width: 160px;
        margin-left: 50px;
    }

    input[type="checkbox"]{
        width: 18px;
        height: 18px;
        margin: 2px;
        margin-left: 8px;
    }

    #register{
        width: 260px;
        height: 40px;
        font-size: 16px;
    }

    #pause{
        font-size: 26px;
        font-weight: bolder;
        color: rgb(38, 190, 7);
        margin:auto;
        padding-left: 6px;
        margin-top: 6px;
    }

    #line{
        background-color: black;
        height: 1px;
        width: 100px;
        display: inline-block;
        position: relative;
        bottom: 7px;
    }

    a{
        font-size: 20px;
        color: rgb(103, 20, 247)
    }
</style>

<Modal {opened} id="modal" target={"body"} on:close={() => (opened = false)}>
    <h1>Create new account</h1>
    
</Modal>

<div id="container">
    <div id="formContainer">
        <button id="register" on:click={() => (opened = true)}>Register</button><br>
        <div id="pause">
            <div id="line"></div>
            OR
            <div id="line"></div>
        </div>
        <form method="post">
            <input type="text" id="login" name="login" placeholder="Login"><br>
            <input type="password" id="password" name="password" placeholder="Password"><br>
            <input type="checkbox" id="if_remember" name="if_remember" value="RememberMe">
            <label for="if_remember" style="font-size: 50%">Remember me</label><br>
            <input type="submit" id="submit" name="submit" value="Submit"><br>
        </form>
    </div>
    <img src={logo} alt="Logo isn't working"><br>
    <a href="">Help, I cannot log in</a><br>
    <a href="../yourservers">Skip</a>
</div>