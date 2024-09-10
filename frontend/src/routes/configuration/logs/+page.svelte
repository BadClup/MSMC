<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "$lib/components/optionsList.svelte";
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

    nav{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    #logsContainer{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 240px);
        float: right;
        flex: 1;
    }

    #logs{
        display: inline-block;
        box-sizing: border-box;
        width: 96%;
        height: 92%;
        margin-top: 2%;
        margin-left: 2%;
        padding: 10px 16px;
        border-radius: 12px 12px 12px 12px;
        background-color: rgb(13, 13, 13);
        border: solid 4px rgb(60, 60, 60);
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav>
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </nav>
        <div id="logsContainer">
            <div id="logs">Logs</div>
        </div>
    </div>
    <Footer/>
</section>