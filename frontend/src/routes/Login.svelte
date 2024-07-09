<script>
</script>
<style>
    #Login{
        width: 100vw;
        font-size: 2em;
        background-color: rgb(88, 88, 88);
        color: rgb(240,240,240);
        padding-left: 20px;
        flex: 1;
    }

    form{
        display: inline-block;
        width: 20%;
        height: 50%;
    }

    img{
        display: inline-block;
        width: 200px;
        height: 200px;
        text-align: center;
    }
</style>

<div id="Login">
    <form method="post" action="">
        <label for="login">Login:</label><br>
        <input type="text" id="login" name="login"><br>
        <label for="password">Password:</label><br>
        <input type="text" id="password" name="password"><br>
        <input type="checkbox" id="if_remember" name="if_remember" value="RememberMe" style="margin: 2px">
        <label for="if_remember" style="font-size: 50%">Remember me</label><br>
        <input type="submit" id="submit" name="submit" value="Sumbit"><br>
    </form>
    <img src="" alt="Logo isn't working">
    <button>Skip</button>
</div>