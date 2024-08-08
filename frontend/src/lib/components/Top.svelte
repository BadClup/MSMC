<script>
    import logo from '$lib/images/logo.svg';
    export let logout = false;
</script>
<style>
    #Top{
        height: 9%;
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
        padding-top: 12px;
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
        display: inline-block;
        width: 30px;
        height: 30px;
        background-size: 30px;
        content: url($lib/images/logout.svg);
        position: relative;
        top: 7px;
        right: 8px;
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