<script lang="ts">
    import plus from '$lib/images/plus.svg';
    import Switch from "../../lib/components/Switch.svelte";
    import { Modal } from '@svelteuidev/core';
    let opened = false;

    interface Server{
        name: string;
        version: string;
        isOnline: boolean;
    }

    const servers: Server[] = [{name: "server1", version:"1.19.2", isOnline:false}, 
    {name: "server2", version:"1.16.4", isOnline:true}, 
    {name: "server3", version:"1.20.1", isOnline:false},
    {name: "server4", version:"1.13.3", isOnline:true}, 
    {name: "server5", version:"1.12.2", isOnline:false}]
</script>

<style>
    section{
        padding-top: 40px;
        text-align: Center;
        font-size: 50px;
        background-color: rgb(88, 88, 88);
        color: rgb(240,240,240);
        flex: 1;
        overflow-y: auto;
    }

    .serverOption{
        display: inline-block;
        height: 100px;
        width: 540px;
        border: dotted black 3px;
        position: relative;
        left: 0px;
        margin-bottom: 15px;
    }

    .serverInfo{
        display: inline-block;
        font-size: 25px;
        height: 100px;
        width: 75%;
        float: left;
        padding-top: 15px;
        padding-left: 20px;
        text-align: left;
    }
    .serverInfo::first-line{
        font-weight: Bold;
    }

    button{
        font-size: 30px; 
        margin-top: 20px;
        background-color: transparent;
        border: 0;
        height: 50px;
        margin-bottom: 50px;
    }

    img{
        font-style:italic; 
        font-size: 15px;
        height: 36px;
        width: 36px;
        position: relative;
        top: 4px;
    }

    p{
        display: inline-block;
        margin: 0;
        position: relative;
        bottom: 2px;
    }

    form{
        position: relative;
        bottom: 15px;
        font-size: 20px;
    }

    h1{
        position: relative;
        bottom: 40px;
        text-align: center;
    }

    input[type="text"], input[type="submit"]{
        box-sizing: border-box;
        width: 100%;
        height: 40px;
        border-radius: 5px;
        padding-left: 10px;
        font-size: 16px;
        margin-top: 10px;
    }

    #submit{
        margin-left: 150px;
        position: relative;
        top: 10px;
        width: 100px;
        height: 40px;
    }

    a{
        width: 100%;
        height: 100%;
        color: inherit;
        text-decoration: none;
    }
</style>

<Modal {opened} id="modal" target={"body"} on:close={() => (opened = false)}>
    <h1>Add a new server</h1>
    <form method="post" action="">
        <label for="name">Server name:</label><br>
        <input type="text" id="name" name="name"><br><br>
        <label for="mcVersion">Server version:</label><br>
        <input list="mcVersionlist" name="mcVersion" id="mcVersion">
            <datalist id="mcVersionlist">
                <option value="1.20">
                <option value="1.19">
            </datalist><br><br>
        <label for="engine">Game engine:</label><br>
        <input list="engineList" name="engine" id="engine">
            <datalist id="engineList">
                <option value="Vanilla">
                <option value="Forge">
                <option value="Fabric">
            </datalist><br><br>
        <label for="engineVersion">Engine version:</label><br>
        <input list="engineVersionlist" name="engineVersion" id="engineVersion">
            <datalist id="engineVersionlist">
                <option value="engineVersion">
                <option value="engineVersion">
                <option value="engineVersion">
            </datalist><br><br>
        <input type="submit" id="submit" name="submit" value="Create">
    </form>
</Modal>

<section>
    <b>Your Servers</b><br><br>
    {#each servers as server}
        <div class="serverOption">
            <div class="serverInfo">
                <a href="../configuration">
                    {server.name}<br>Server version: {server.version}
                </a>
            </div>
            <div style="float: right; position: relative; top: 33px; right: 15px">
                <Switch state={server.isOnline}/>
            </div>
        </div><br>
    {/each}
    <button on:click={() => (opened = true)}>
        <img src={plus} alt="Cannot load a graphic">
        <p>Add a server<p>
    </button>
</section>