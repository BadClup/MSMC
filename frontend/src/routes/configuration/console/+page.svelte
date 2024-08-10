<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "./../optionsList.svelte";
</script>
<style>
    section{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    section > div{
        width: 100%;
        flex: 1;
    }

    #listContainer{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    #consoleContainer{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 240px);
        float: right;
        padding-top: 15px;
        flex: 1;
    }

    #consoleHistory{
        display: inline-block;
        box-sizing: border-box;
        width: 90%;
        height: 90%;
        border-radius: 10px 10px 0 0;
        margin-left: 5%;
        background-color: black;
    }

    input{
        display: inline-block;
        box-sizing: border-box;
        width: 90%;
        height: 40px;
        border-radius: 0 0 10px 10px;
        margin-left: 5%;
        font-size: 16px;
        padding-left: 10px;
    }
    
    input:focus{
        outline: none;
    }

</style>

<section>
    <Top showServers={true} canLogout={true}/>
    <div>
        <div id="listContainer">
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </div>
        <div id="consoleContainer">
            <div id="consoleHistory"></div>
            <input type="text" autofocus spellcheck="false">
        </div>
    </div>
    <Footer/>
</section>