"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import grpc
import prodvana.proto.prodvana.recipe.manager_pb2

class RecipeManagerStub:
    def __init__(self, channel: grpc.Channel) -> None: ...
    ConfigureRecipe: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeReq,
        prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeResp,
    ]
    ValidateConfigureRecipe: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeReq,
        prodvana.proto.prodvana.recipe.manager_pb2.ValidateConfigureRecipeResp,
    ]
    ListRecipes: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.recipe.manager_pb2.ListRecipesReq,
        prodvana.proto.prodvana.recipe.manager_pb2.ListRecipesResp,
    ]
    GetRecipe: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeReq,
        prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeResp,
    ]
    GetRecipeConfig: grpc.UnaryUnaryMultiCallable[
        prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeConfigReq,
        prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeConfigResp,
    ]

class RecipeManagerServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def ConfigureRecipe(
        self,
        request: prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeResp: ...
    @abc.abstractmethod
    def ValidateConfigureRecipe(
        self,
        request: prodvana.proto.prodvana.recipe.manager_pb2.ConfigureRecipeReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.recipe.manager_pb2.ValidateConfigureRecipeResp: ...
    @abc.abstractmethod
    def ListRecipes(
        self,
        request: prodvana.proto.prodvana.recipe.manager_pb2.ListRecipesReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.recipe.manager_pb2.ListRecipesResp: ...
    @abc.abstractmethod
    def GetRecipe(
        self,
        request: prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeResp: ...
    @abc.abstractmethod
    def GetRecipeConfig(
        self,
        request: prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeConfigReq,
        context: grpc.ServicerContext,
    ) -> prodvana.proto.prodvana.recipe.manager_pb2.GetRecipeConfigResp: ...

def add_RecipeManagerServicer_to_server(servicer: RecipeManagerServicer, server: grpc.Server) -> None: ...