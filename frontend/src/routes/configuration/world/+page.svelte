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

    main{
        display: inline-block;
        box-sizing: border-box;
        width: calc(100% - 240px);
        height: 100%;
        padding-top: 2%;
    }

    h1{
        text-align: center;
        font-weight: bold;
        font-size: 60px;
        color: inherit;
        margin: 0;
        margin-bottom: 40px;
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav>
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </nav>
        <main>
            <h1>World</h1>
        </main>
    </div>
    <Footer/>
</section>