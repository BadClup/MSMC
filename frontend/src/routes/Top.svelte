<script>
    import logo from '$lib/images/logo.svg';
    import out_img from '$lib/images/logout.png';
    export let logout = false;
</script>
<style>
    #Top{
        height: 8%;
        width: 100vw;
        box-sizing: border-box;
        display: flex;
    }

    #left{
        width: 45%;
        height: 100%;
        text-align: left;
        font-size: 3em;
        font-weight: bolder;
        background-color: rgb(40,40,40);
        color: rgb(240,240,240);
        padding-top: 8px;
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
        padding-top: 13px;
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
        top: 8px;
    }

    button{
        width: 120px;
        height: 30px;
        font-size: 20px;
        font-weight: bold;
        color: white;
        background-color: transparent;
        border: 0;
    }

    #logout::before{
        content: url(../lib/images/logout.png);
        position: relative;
        top: 9px;
        right: 4px;
    }
</style>

<div id="Top">
    <div id="left">
        <img src={logo} alt="Logo isn't working">
        MSMC
    </div>

    <div id="right">
        {#if logout == true}
            <button id="logout">Log out</button>
        {/if}
    </div>
</div>