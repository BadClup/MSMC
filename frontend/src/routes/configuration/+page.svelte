<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "./optionsList.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    #container2{
        width: 100%;
        flex: 1;
    }

    #listContainer{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    #settings{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 240px);
        float: right;
        padding-top: 15px;
        flex: 1;
    }

    h1{
        position: relative;
        text-align: center;
        font-weight: bold;
        font-size: 50px;
        color: inherit;
    }

    .infoContainer{
        margin-left: 60px;
    }
    
    .label{
        display: inline-block;
        font-size: 25px;
        color: inherit;
        border-radius: 3px;
        background-color: rgb(32, 129, 226);
        padding: 3px;
        padding-bottom: 0;
    }
    
    .info{
        display: inline-block;
        font-size: 25px;
        color: black;
        border-radius: 2px;
        background-color: rgb(231, 231, 231);
        padding: 3px;
        padding-bottom: 0;
        width: 340px;
    }
</style>

<div id="container">
    <Top logout={true}/>
    <div id="container2">
        <div id="listContainer">
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </div>
        <div id="settings">
            <h1>Server name</h1><br>
            <div class="infoContainer">
                <div class="label">Engine</div><br>
                <div class="info"> Engine info</div>
            </div><br><br>
            <div class="infoContainer">
                <div class="label">Game version</div><br>
                <div class="info"> Version info</div>
            </div>
        </div>
    </div>
    <Footer/>
</div>