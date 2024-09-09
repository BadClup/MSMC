<script>
    import logo from '$lib/images/logo.svg';
    export let isLoggedIn = false;
    export let showServers = false;
</script>
<style>
    header{
        height: 85px;
        width: 100vw;
        box-sizing: border-box;
        display: flex;
    }

    #left{
        width: 45%;
        height: 100%;
        text-align: left;
        font-size: 48px;
        font-weight: bolder;
        background-color: rgb(40,40,40);
        color: rgb(240,240,240);
        padding-top: 4px;
        padding-right: 40px;
        padding-left: 30px;
        box-sizing: border-box;
        flex: 1;
    }

    #right{
        width: 45%;
        height: 100%;
        text-align: right;
        font-size: 1em;
        background-color: rgb(40,40,40);
        color: rgb(240,240,240);
        padding-top: 16px;
        padding-top: 16px;
        padding-right: 40px;
        padding-left: 30px;
        box-sizing: border-box;
        flex: 1;
    }

    img{
        font-size: 8px;
        width: 50px;
        height: 50px;
        position: relative;
        top: 11px;
    }

    #logout, #servers{
        width: 120px;
        height: 30px;
        font-size: 20px;
        font-weight: bold;
        color: white;
        background-color: transparent;
        border: 0;
    }

    #servers{
        margin-right: 20px;
    }

    #logout::before{
        display: inline-block;
        width: 30px;
        height: 30px;
        background-size: 30px;
        content: url($lib/images/logout.svg);
        position: relative;
        top: 7px;
        right: 8px;
    }

    #servers::before{
        display: inline-block;
        width: 30px;
        height: 30px;
        background-size: 30px;
        content: url($lib/images/servers.svg);
        position: relative;
        top: 7px;
        right: 8px;
    }

    @media screen and (max-width: 600px) {
        #left{
            font-size: 32px;
        }
    }
</style>

<header>
    <div id="left">
        <img src={logo} alt="Logo isn't working">
        MSMC
    </div>

    <div id="right">
        {#if showServers == true}
            <button id="servers">Servers</button>
        {/if}
        {#if isLoggedIn == true}
            <button id="logout">Log out</button>
        {/if}
    </div>
</header>