<script>
    import logo from '$lib/images/logo.svg';
    import { Modal } from '@svelteuidev/core';
    let opened = false;
</script>
<style>
    section{
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
        position: relative;
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

    #submit{
        background-color: rgb(82, 192, 67);
        width: 200px;
        margin-left: 30px;
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
        color: rgb(34, 196, 2);
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

    h1{
        margin: auto;
        text-align: center;
        position: relative;
        bottom: 35px;
        bottom: 15px;
    }

    #modalForm{
        display: inline-block;
        width: 88%;
        height: 80%;
        margin-left: 38px;
    }

    #newLogin, #newPassword, #newPassword2{
        box-sizing: border-box;
        width: 90%;
        height: 40px;
        border-radius: 5px;
        padding-left: 10xp;
        font-size: 16px;
        margin-top: 10px;
    }

    #create{
        width: 200px;
        height: 50px;
        margin-left: 60px;
        margin-top: 30px;
    }
</style>

<Modal {opened} id="modal" target={"body"} on:close={() => (opened = false)}>
    <h1>Create new account</h1>
    <form method="post" action="" id="modalForm">
        <input type="text" id="newLogin" name="newlogin" placeholder="Enter your login">
        <input type="password" id="newPassword" name="newpassword" placeholder="Enter your password">
        <input type="password" id="newPassword2" name="newpassword2" placeholder="Repeat your password">
        <input type="submit" id="create" name="create" value="Create">
    </form>
</Modal>

<section>
    <div id="formContainer">
        <button id="register" on:click={() => (opened = true)}>Register</button>
        <div id="pause">
            <div id="line"></div>
            OR
            <div id="line"></div>
        </div>
        <form method="post">
            <input type="text" id="login" name="login" placeholder="Login">
            <input type="password" id="password" name="password" placeholder="Password">
            <input type="checkbox" id="if_remember" name="if_remember" value="RememberMe">
            <label for="if_remember" style="font-size: 50%">Remember me</label>
            <input type="submit" id="submit" name="submit" value="Submit">
        </form>
    </div>
    <img src={logo} alt="Logo isn't working"><br>
    <a href="#">Help, I cannot log in</a><br>
    <a href="../servers">Skip</a>
</section>