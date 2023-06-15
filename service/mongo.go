package service

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var globalClient *mongo.Client

func ConnectDb() error {
	if globalClient != nil {
		// Já existe uma conexão estabelecida
		return nil
	}
	connStr := os.Getenv("URI")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.
		Client().
		ApplyURI(connStr).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	globalClient = client
	return nil
}

func FindClassByStudent(alunoID string) (bson.M, error) {
	collectionStudent := globalClient.Database("banco1").Collection("usuarios")
	collectionClass := globalClient.Database("banco2").Collection("turma-example")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Consulta para buscar o aluno pelo ID
    alunoFilter := bson.M{"_id": alunoID}
    var aluno bson.M
    if err := collectionStudent.FindOne(ctx, alunoFilter).Decode(&aluno); err != nil {
        return nil, err
    }

    // Consulta para buscar a turma do aluno
    turmaFilter := bson.M{"id_aluno": alunoID}
	turmaProjection := bson.M{"turma": 1, "turno": 1}
    var turma bson.M
    if err := collectionClass.FindOne(ctx, turmaFilter, options.FindOne().SetProjection(turmaProjection)).Decode(&turma); err != nil {
        return nil, err
    }

	aluno["turma"] = turma

	return aluno, nil
}

//Exemplo para duas connection string
func ConnectWithOnlyDB(alunoID string) (bson.M, error) {
	// Conexão com o primeiro banco de dados
	connStrDb1 := os.Getenv("URI_DB1")
	connStrDb2 := os.Getenv("URI_DB2")
	client1, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStrDb1))
	if err != nil {
		log.Fatal(err)
	}
	defer client1.Disconnect(context.Background())

	// Conexão com o segundo banco de dados
	client2, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStrDb2))
	if err != nil {
		log.Fatal(err)
	}
	defer client2.Disconnect(context.Background())

	// Exemplo de uso das conexões
	db1 := client1.Database("banco1").Collection("usuarios")
	db2 := client1.Database("banco2").Collection("turma-example")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Consulta para buscar o aluno pelo ID
    alunoFilter := bson.M{"_id": alunoID}
    var aluno bson.M
    if err := db1.FindOne(ctx, alunoFilter).Decode(&aluno); err != nil {
        return nil, err
    }

    // Consulta para buscar a turma do aluno
    turmaFilter := bson.M{"id_aluno": alunoID}
	turmaProjection := bson.M{"turma": 1, "turno": 1}
    var turma bson.M
    if err := db2.FindOne(ctx, turmaFilter, options.FindOne().SetProjection(turmaProjection)).Decode(&turma); err != nil {
        return nil, err
    }

	aluno["turma"] = turma

	return aluno, nil
}